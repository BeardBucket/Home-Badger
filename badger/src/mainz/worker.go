package mainz

import (
	"errors"
	"github.com/BeardBucket/Home-Badger/src/hasser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ErrNoVprGiven        = errors.New("viper passed in is nil")
	ErrNoCmdGiven        = errors.New("command passed in is nil")
	ErrMainWorkerDefined = errors.New("main (worker) already defined")
)

var main Main
var worker MainWorker

type Main interface {
	onLateInit() error
	onRun() error
	onExit() error
	onCycle() error
	Vpr() *viper.Viper
	FailIt(msg string, err error)
}

type MainWorker struct {
	vpr          *viper.Viper   // Config
	cmd          *cobra.Command // Command that was run - should be "Run"
	cmdArgs      []string       // Any positional args
	notifyF      NotifyF        // Notify of a fatal error and exit
	logger       *logrus.Logger // Primary Logger
	logLevelText string         // Text log level passed in
	optionsPath  string         // Path to HASS Add-On options file (JSON)
}

func NewMainWorker(cmd *cobra.Command, args []string, notifyF NotifyF, vpr *viper.Viper, hassAddOnOptionsPath string, logLevel string) (*MainWorker, error) {
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
	if logLevel == "" {
		logLevel = "info"
	}

	worker = MainWorker{
		vpr:          vpr,
		cmd:          cmd,
		cmdArgs:      args,
		notifyF:      notifyF,
		logger:       logrus.New(),
		optionsPath:  hassAddOnOptionsPath,
		logLevelText: logLevel,
	}
	err := worker.setupLogging()
	if err != nil {
		return nil, err
	}

	main = worker

	return &worker, nil
}

func GetMain() Main {
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

// onCycle should be called occasionally by the main thread
func (w MainWorker) onCycle() error {
	w.L().Debug("hai!")
	return nil
}

// onRun should be called once to start the main process(es)
func (w MainWorker) onRun() error {
	w.L().Debug("runnnnnnnnnnn")
	go func() {
		ehass, _ := hasser.NewEventHass()
		err := ehass.Test()
		if err != nil {
			w.FailIt("Problem creating EventHass", err)
		}
	}()
	return nil
}

// onLateInit should be called once just before onRun() is called
func (w MainWorker) onLateInit() error {
	w.L().Debug("late init")

	return nil
}

// onExit should be called when an early exit is required
func (w MainWorker) onExit() error {
	w.L().Debug("cleanup")
	return nil
}
