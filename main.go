package main

import (
	"api-gateway-golang/server"
	"api-gateway-golang/server/configs/env"
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/middlewares"

	"api-gateway-golang/server/routers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	goEnv := env.GetEnv("GO_ENV")

	if goEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InitLogger()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.SystemLog.Debug(httpMethod + ": " + absolutePath + " --> " + handlerName + " (" + strconv.Itoa(nuHandlers) + " handlers)")
	}

	router := server.CreateRouter()
	router.NoRoute(middlewares.NoRouteHandle)
	router.Use(middlewares.ExceptionHandle)
	router.Use(middlewares.AcceptRequestHandle)
	router.Use(middlewares.WhitePrefixHandle)

	routers.RegisterRoutes(router)

	logger.SystemLog.Info("Server is running on port " + env.GetEnv("PORT"))
	http.ListenAndServe(":"+env.GetEnv("PORT"), router)
}
