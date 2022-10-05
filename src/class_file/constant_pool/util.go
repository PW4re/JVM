package constant_pool

import "log"

var logger = func() *log.Logger {
	logger := log.Default()
	logger.SetPrefix("[CP_Parser]")
	return logger
}()
