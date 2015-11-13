package main

import (
	"github.com/zlbbq/go-logger"
)

type Struct struct {
	Name string
	Age int
}

func main() {
	logger.Debug("Hello, %s", "World")
	logger.Info("Hello, %s", "World")
	logger.Warn("Hello, %s", "World")
	logger.ErrorD("Hello, %s", "World")
	logger.FatalD("Hello, %s", "World")

	// It is strongly recommended that call logger.Get() function to get a logger from logger pool
	moduleLogger := logger.Get("module_full_name")
	moduleLogger.Info("This is a module logger, it is able to be controlled by application entry")

	s := Struct{
		"Zlbbq", 20,
	}

	logger.DebugLog(s)
}
