package services

import (
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/configs/response"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type otherAuthService struct{}

var OtherAuthService = otherAuthService{}

func (svc otherAuthService) OtherAuth(c *gin.Context) {
	path := c.Param("path")

	headerJson, _ := json.MarshalIndent(c.Request.Header, "", "    ")
	logger.Log.Info(string(headerJson))

	response.Ctx(c).Success(map[string]string{"path": path, "auth": "other"})
}
