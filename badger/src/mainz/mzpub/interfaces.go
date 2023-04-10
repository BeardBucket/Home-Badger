package mzpub

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Main interface {
	OnLateInit() error
	OnRun() error
	OnExit() error
	OnCycle() error
	Vpr() *viper.Viper
	FailIt(msg string, err error)
	L() *logrus.Logger
}
