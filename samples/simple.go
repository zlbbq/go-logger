package main

import (
	"github.com/zlbbq/go-logger"
)

func main() {
	logger.Debug("Hello, %s", "World")
	logger.Info("Hello, %s", "World")
	logger.Warn("Hello, %s", "World")
	logger.Error("Hello, %s", "World")
	logger.Fatal("Hello, %s", "World")

	// It is strongly recommended that call logger.Get() function to get a logger from logger pool
	moduleLogger := logger.Get("module_full_name")
	moduleLogger.Info("This is a module logger, it is able to be controlled by application entry")
}
