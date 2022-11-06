package class_file

import "log"

var logger = func() *log.Logger {
	logger := log.Default()
	logger.SetPrefix("[Class]")
	return logger
}()
