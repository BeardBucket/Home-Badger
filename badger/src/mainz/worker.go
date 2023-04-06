package mainz

import (
	"errors"
	"fmt"
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
}

type MainWorker struct {
	vpr     *viper.Viper   // Config
	cmd     *cobra.Command // Command that was run - should be "Run"
	cmdArgs []string       // Any positional args
	notifyF NotifyF        // Notify of a fatal error and exit
	logger  *logrus.Logger // Primary Logger
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

	worker = MainWorker{
		vpr:     vpr,
		cmd:     cmd,
		cmdArgs: args,
		notifyF: notifyF,
		logger:  logrus.New(),
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

// onCycle should be called occasionally by the main thread
func (w MainWorker) onCycle() error {
	return nil
}

// onRun should be called once to start the main process(es)
func (w MainWorker) onRun() error {
	return nil
}

// onLateInit should be called once just before onRun() is called
func (w MainWorker) onLateInit() error {
	return nil
}

// onExit should be called when an early exit is required
func (w MainWorker) onExit() error {
	fmt.Println("cleanup")
	return nil
}
