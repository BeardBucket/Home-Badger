package mainz

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type (
	NotifyF func(msg string, err error)
)

// OnLateInit is called after basic startup
func OnLateInit(cmd *cobra.Command, args []string, notifyF NotifyF, vpr *viper.Viper) error {
	worker, err := NewMainWorker(cmd, args, notifyF, vpr)
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
