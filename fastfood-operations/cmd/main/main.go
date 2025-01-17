package main

import (
	logger "github.com/sirupsen/logrus"
)

// @title Fast-food Operations
// @version 1.0.0
// @description The Fast-food Operations.
// @contact.email postec8soatg46@gmail.com
// @licence.name Copyright (c) 2022. All rights reserved.
func main() {
	logger.Info("Application execution started...")

	defer handlePanic()

	logger.Info("Applying CORS Policies for Local Environment")

	engine := GetRouters(true)

	logger.Info("Going to initialize the server.")
	panic(engine.Run(":8080"))
}

func handlePanic() {
	if r := recover(); r != nil {
		logger.WithField("panic", r).Error("Panic occurred in the application.")
	}
}
