package main

import (
	"os"

	"oss.nandlabs.io/golly/l3"
	"oss.nandlabs.io/golly/lifecycle"
)

var manager = lifecycle.NewSimpleComponentManager()
var logger = l3.Get()

func main() {
	// Determine the application environment (default to "local" if not set)
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}
	logger.InfoF("Starting application in %s environment", appEnv)
	// Build the configuration file name based on the environment
	configFileName := "config-" + appEnv + ".yml"
	logger.InfoF("Using configuration file: %s", configFileName)
	// Load the configuration file
	_, err := loadConfig(configFileName)
	if err != nil {
		logger.ErrorF("Error loading config:", err)
		return
	}
	start()
	defer shutdown()
}

func start() {
	restServer := getRestServer(manager)
	manager.Register(restServer)
	manager.StartAndWait()
}

func shutdown() {
	manager.StopAll()
}
