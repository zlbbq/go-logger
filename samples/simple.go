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
}
