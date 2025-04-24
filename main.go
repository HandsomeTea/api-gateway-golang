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
	goEnv, _ := env.GetEnv("GO_ENV")

	if goEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	defer logger.Close()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.SystemLog.Debug(httpMethod + ": " + absolutePath + " --> " + handlerName + " (" + strconv.Itoa(nuHandlers) + " handlers)")
	}

	router := server.CreateRouter()
	router.NoRoute(middlewares.NoRouteHandle)
	router.Use(middlewares.ExceptionHandle)
	router.Use(middlewares.AcceptRequestHandle)
	router.Use(middlewares.WhitePrefixHandle)

	routers.RegisterRoutes(router)

	port, _ := env.GetEnv("PORT")

	logger.SystemLog.Info("Server is running on port " + port)
	http.ListenAndServe(":"+port, router)
}
