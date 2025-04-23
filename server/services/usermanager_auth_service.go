package services

import (
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/configs/response"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type userManagerAuthService struct{}

var UserManagerAuthService = userManagerAuthService{}

func (svc userManagerAuthService) AdminAuth(c *gin.Context) {
	path := c.Param("path")

	// print(path[100])

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "客户端验证"})
}

func (svc userManagerAuthService) LoginAuth(c *gin.Context) {
	path := c.Param("path")

	headerJson, _ := json.MarshalIndent(c.Request.Header, "", "    ")
	logger.Log.Info(string(headerJson))

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "需要登录验证"})
}
