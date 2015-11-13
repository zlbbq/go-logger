package main

import (
	goLogger "github.com/zlbbq/go-logger"
	"fmt"
	"os"
)

func main() {
	// Get a logger form logger pool
	logger := goLogger.Get("logger-pool")
	// Common usage
	logger.Debug("Hello, %s", "World")
	logger.Info("Hello, %s", "World")
	logger.Warn("Hello, %s", "World")
	logger.ErrorD("Hello, %s", "World")
	logger.FatalD("Hello, %s", "World")
	fmt.Println("=====================================================================================")
	// Change log level
	logger.Level = goLogger.LevelInfo
	logger.Debug("Hello, %s", "World")			// Disappeared
	logger.Info("Hello, %s", "World")
	logger.Warn("Hello, %s", "World")
	logger.ErrorD("Hello, %s", "World")
	logger.FatalD("Hello, %s", "World")
	fmt.Println("=====================================================================================")
	// Disable colorful
	logger.Colorful = false
	logger.Debug("Hello, %s", "World")			// Disappered too
	logger.Info("Hello, %s", "World")
	logger.Warn("Hello, %s", "World")
	logger.ErrorD("Hello, %s", "World")
	logger.FatalD("Hello, %s", "World")
	fmt.Println("=====================================================================================")
	// Create and register a logger to logger pool
	nl := goLogger.NewLogger("github.com/xxx/xxx", goLogger.LevelError, true, os.Stdout)
	goLogger.Register(nl)
	nl.Debug("Hello, %s", "World")				// Disappeared
	nl.Info("Hello, %s", "World")					// Disappeared
	nl.Warn("Hello, %s", "World")					// Disappeared
	nl.ErrorD("Hello, %s", "World")
	nl.FatalD("Hello, %s", "World")
	fmt.Println("=====================================================================================")
	// A module get logger from logger pool
	moduleLogger := goLogger.Get("github.com/xxx/xxx")
	moduleLogger.Debug("Hello, %s", "World")				// Disappeared too
	moduleLogger.Info("Hello, %s", "World")					// Disappeared too
	moduleLogger.Warn("Hello, %s", "World")					// Disappeared too
	moduleLogger.ErrorD("Hello, %s", "World")
	moduleLogger.FatalD("Hello, %s", "World")
	fmt.Println("=====================================================================================")
	// Redirect output
	logger.Level = goLogger.LevelDebug
	logger.Colorful = true
	logger.Output = os.Stdout								// A demostration
	logger.Debug("Hello, %s", "World")
	logger.Info("Hello, %s", "World")
	logger.Warn("Hello, %s", "World")
	logger.ErrorD("Hello, %s", "World")
	logger.FatalD("Hello, %s", "World")
}
