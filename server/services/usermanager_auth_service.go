package services

import (
	"api-gateway-golang/server/configs/response"

	"github.com/gin-gonic/gin"
)

type userManagerAuthService struct{}

var UserManagerAuthService = userManagerAuthService{}

func (svc userManagerAuthService) authenticate(path string, auth string) (bool, map[string]string) {

	// panic(globals.NewException("other auth service not implemented"))
	return true, map[string]string{"path": path, "auth": auth}
}

// ============================================ gin接口处理函数 ============================================

func (svc userManagerAuthService) AdminAuth(c *gin.Context) {
	path := c.Param("path")
	ok, data := svc.authenticate(path, "admin验证")

	if !ok {
		response.Ctx(c).Failed("Authentication failed")
		return
	}

	response.Ctx(c).Success(data)
}

func (svc userManagerAuthService) LoginAuth(c *gin.Context) {
	path := c.Param("path")
	ok, data := svc.authenticate(path, "需要登录验证")

	if !ok {
		response.Ctx(c).Failed("Authentication failed")
		return
	}

	response.Ctx(c).Success(data)
}
