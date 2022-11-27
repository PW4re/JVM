package cp

import "log"

var logger = func() *log.Logger {
	logger := log.Default()
	logger.SetPrefix("[CP]")
	return logger
}()
