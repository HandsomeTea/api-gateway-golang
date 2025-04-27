package services

import (
	"api-gateway-golang/server/configs/response"

	"github.com/gin-gonic/gin"
)

type userManagerAuthService struct{}

var UserManagerAuthService = userManagerAuthService{}

func (svc userManagerAuthService) AuthenticateService(path string, auth string) map[string]string {

	// panic(globals.NewException("other auth service not implemented"))
	return map[string]string{"path": path, "auth": auth}
}

// ============================================ gin接口处理函数 ============================================

func (svc userManagerAuthService) AdminAuth(c *gin.Context) {
	path := c.Param("path")
	data := svc.AuthenticateService(path, "admin验证")

	response.Ctx(c).Success(data)
}

func (svc userManagerAuthService) LoginAuth(c *gin.Context) {
	path := c.Param("path")
	data := svc.AuthenticateService(path, "需要登录验证")

	response.Ctx(c).Success(data)
}
