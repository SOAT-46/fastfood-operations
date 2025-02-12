package main

import (
	"net/http"

	_ "github.com/SOAT-46/fastfood-operations/cmd/docs" // for swagger
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// CorsMiddleware CORS config for App access.
func CorsMiddleware() gin.HandlerFunc {
	return func(cors *gin.Context) {
		logger.Info("Setting up CORS Middleware")

		cors.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		cors.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		cors.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if cors.Request.Method == http.MethodOptions {
			cors.Writer.WriteHeader(http.StatusOK)
			return
		}
		cors.Next()
	}
}

func GetRoutes(corsFlag bool) *gin.Engine {
	logger.Info("Setting up Routes")

	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if corsFlag {
		engine.Use(CorsMiddleware())
	}

	groups := map[string]*gin.RouterGroup{}

	apps := injectApps()
	for _, app := range apps {
		for _, controller := range app.GetControllers() {
			bind := controller.GetBind()

			group, found := groups[bind.Version]
			if !found {
				group = engine.Group(bind.Version)
				groups[bind.Version] = group
			}
			group.Handle(bind.Method, bind.RelativePath, controller.Execute)
		}
	}

	logger.Info("Routes setup completed")

	return engine
}
