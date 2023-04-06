package mainz

// OnLateInit is called after basic startup
func OnLateInit() error {
	worker, err := NewMainWorker()
	if err != nil {
		return err
	}
	main = worker
	err = main.onLateInit()
	if err != nil {
		return err
	}
	return nil
}
