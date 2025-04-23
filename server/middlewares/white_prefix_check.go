package middlewares

import (
	"api-gateway-golang/server/configs"
	httpError "api-gateway-golang/server/configs/error"
	"api-gateway-golang/server/configs/response"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func WhitePrefixHandle(c *gin.Context) {
	if !slices.ContainsFunc(configs.GatewayConfig.WhitePrefixes, func(s string) bool {
		return strings.HasPrefix(c.Request.URL.Path, s)
	}) {
		response.Ctx(c).Failed("unknown request", httpError.FORBIDDEN)
		c.Abort()
		return
	}
	c.Next()
}
