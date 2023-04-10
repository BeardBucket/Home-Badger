package mainz

func OnCycle() error {
	err := GetMain().OnCycle()
	if err != nil {
		return err
	}
	return nil
}
