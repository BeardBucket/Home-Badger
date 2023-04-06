package mainz

import "fmt"

var main Main

type Main interface {
	onLateInit() error
	onRun() error
	onExit() error
}

type MainWorker struct {
}

func NewMainWorker() (Main, error) {
	w := MainWorker{}

	return &w, nil
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
