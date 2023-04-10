package mainz

import (
	"errors"
	"github.com/BeardBucket/Home-Badger/src/hasser"
	"github.com/BeardBucket/Home-Badger/src/mainz/mzpub"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ErrNoVprGiven        = errors.New("viper passed in is nil")
	ErrNoCmdGiven        = errors.New("command passed in is nil")
	ErrMainWorkerDefined = errors.New("main (worker) already defined")
)

var main mzpub.Main
var worker MainWorker

type MainWorker struct {
	vpr           *viper.Viper         // Config
	cmd           *cobra.Command       // Command that was run - should be "Run"
	cmdArgs       []string             // Any positional args
	notifyF       NotifyF              // Notify of a fatal error and exit
	logger        *logrus.Logger       // Primary Logger
	logLevelText  string               // Text log level passed in
	optionsPath   string               // Path to HASS Add-On options file (JSON)
	webListenPort int                  // What port our webserver should listen on
	cache         *cache.Cache[string] // Main cache
}

func NewMainWorker(cmd *cobra.Command, args []string, notifyF NotifyF, vpr *viper.Viper) (*MainWorker, error) {
	// Gotta have a non-nil viper
	if vpr == nil {
		return nil, ErrNoVprGiven
	}
	// Gotta have a non-nil command
	if cmd == nil {
		return nil, ErrNoCmdGiven
	}
	if main != nil {
		return nil, ErrMainWorkerDefined
	}

	logLevel, err := cmd.Flags().GetString("level")
	if err != nil {
		return nil, err
	}
	if logLevel == "" {
		logLevel = "info"
	}

	hassAddOnOptionsPath, err := cmd.Flags().GetString("addon-options")
	if err != nil {
		return nil, err
	}

	listenPort, err := cmd.Flags().GetInt("port")
	if err != nil {
		return nil, err
	}

	worker = MainWorker{
		vpr:           vpr,
		cmd:           cmd,
		cmdArgs:       args,
		notifyF:       notifyF,
		logger:        logrus.New(),
		optionsPath:   hassAddOnOptionsPath,
		logLevelText:  logLevel,
		webListenPort: listenPort,
	}
	if err := worker.setupLogging(); err != nil {
		return nil, err
	}

	main = worker

	return &worker, nil
}

func GetMain() mzpub.Main {
	return main
}

func (w MainWorker) Vpr() *viper.Viper {
	return w.vpr
}

// FailIt gracefully exits the program due to a unrecoverable error of some kind
func (w MainWorker) FailIt(msg string, err error) {
	w.L().WithError(err).Error(msg)
	w.notifyF(msg, err)
}

// OnCycle should be called occasionally by the main thread
func (w MainWorker) OnCycle() error {
	w.L().Debug("hai!")
	return nil
}

// OnRun should be called once to start the main process(es)
func (w MainWorker) OnRun() error {
	w.L().Debug("runnnnnnnnnnn")
	go func() {
		hass, _ := hasser.NewHass(w)
		err := hass.TestingF()
		if err != nil {
			w.FailIt("Problem creating Hass", err)
		}
	}()
	return nil
}

// OnExit should be called when an early exit is required
func (w MainWorker) OnExit() error {
	w.L().Debug("cleanup")
	return nil
}
