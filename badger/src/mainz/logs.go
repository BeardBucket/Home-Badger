package mainz

import (
	"github.com/sirupsen/logrus"
	"os"
)

func (w MainWorker) setupLogging() error {
	w.L().SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		//ForceQuote:                false,
		//DisableQuote:              false,
		//EnvironmentOverrideColors: false,
		//DisableTimestamp:          false,
		FullTimestamp: true,
		//TimestampFormat:           "",
		//DisableSorting:            false,
		//SortingFunc:               nil,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		QuoteEmptyFields:       true,
		//FieldMap:                  nil,
		//CallerPrettyfier:          nil,
	})
	level, err := logrus.ParseLevel(w.logLevelText)
	if err != nil {
		return err
	}
	w.L().SetLevel(level)
	w.L().Out = os.Stdout

	return nil
}

func (w MainWorker) L() *logrus.Logger {
	return w.logger
}
