package middlewares

import (
	"api-gateway-golang/server/configs"
	httpError "api-gateway-golang/server/configs/error"
	"api-gateway-golang/server/configs/response"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func WhitePrefixHandle(c *gin.Context) {
	whitePrefixes := configs.GetGatewayConfig().WhitePrefixes
	normalizedPath := strings.TrimRight(c.Request.URL.Path, "/")
	cleanPath := path.Clean(normalizedPath)

	if !strings.HasPrefix(cleanPath, "/") {
		cleanPath = "/" + cleanPath // 确保绝对路径
	}
	match := false

	for _, prefix := range whitePrefixes {
		if strings.HasPrefix(cleanPath, prefix) {
			match = true
			break
		}
	}

	if !match {
		response.Ctx(c).Failed("access denied", httpError.FORBIDDEN)
		c.Abort()
		return
	}
	c.Next()
}
