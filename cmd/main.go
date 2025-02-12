package main

import (
	logger "github.com/sirupsen/logrus"
)

// @title Fastfood Operations
// @version 1.0.0
// @description Fastfood Operations.
// @contact.email postec8soatg46@gmail.com
// @licence.name Copyright (c) 2025. All rights reserved.
func main() {
	logger.Info("Application execution started...")
	defer handlePanic()

	logger.Info("Applying CORS Policies")
	engine := GetRoutes(true)

	logger.Info("Going to initialize the server.")
	panic(engine.Run(":8080"))
}

func handlePanic() {
	if r := recover(); r != nil {
		logger.WithField("panic", r).Error("Panic occurred in the application.")
	}
}
