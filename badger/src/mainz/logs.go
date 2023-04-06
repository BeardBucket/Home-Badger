package mainz

import log "github.com/sirupsen/logrus"

func (w MainWorker) L() *log.Logger {
	return w.logger
}
