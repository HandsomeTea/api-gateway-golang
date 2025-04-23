package middlewares

import (
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/configs/response"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ExceptionHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%v", err)
			response.Ctx(c).Failed(message)
			logger.SystemLog.Error(message + "\n" + string(debug.Stack()))
		}
	}()

	c.Next()
}

func NoRouteHandle(c *gin.Context) {
	response.Ctx(c).Failed("404 not found", "URL_NOT_FOUND", c.Request.Method+": "+c.Request.URL.RequestURI())
	c.Abort()
}
