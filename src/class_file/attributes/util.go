package attributes

import "log"

var logger = func() *log.Logger {
	logger := log.Default()
	logger.SetPrefix("[Attributes]")
	return logger
}()
