package middlewares

import (
	httpError "api-gateway-golang/server/configs/error"
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/configs/response"
	"fmt"
	"html"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ExceptionHandle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%v", err)
			logger.SystemLog.Error(message + "\n" + string(debug.Stack()))
			response.Ctx(c).Failed(message)
			c.Abort()
		}
	}()

	c.Next()
}

func NoRouteHandle(c *gin.Context) {
	response.Ctx(c).Failed("url not found", httpError.URL_NOT_FOUND, c.Request.Method+": "+html.EscapeString(c.Request.URL.RequestURI()))
	c.Abort()
}
