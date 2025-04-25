package services

import (
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/configs/response"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type otherAuthService struct{}

var OtherAuthService = otherAuthService{}

func (svc otherAuthService) AuthenticateService(path string, headers map[string][]string) (bool, map[string]string) {
	headerJson, _ := json.MarshalIndent(headers, "", "    ")
	logger.Log.Info(string(headerJson))

	// panic(globals.NewException("other auth service not implemented"))
	return true, map[string]string{"path": path, "auth": "other"}
}

// ============================================ gin接口处理函数 ============================================
func (svc otherAuthService) OtherAuth(c *gin.Context) {
	path := c.Param("path")
	ok, data := svc.AuthenticateService(path, c.Request.Header)

	if !ok {
		response.Ctx(c).Failed("Authentication failed")
		return
	}

	response.Ctx(c).Success(data)
}
