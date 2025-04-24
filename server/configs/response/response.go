package response

import (
	"encoding/json"
	"net/http"

	httpError "api-gateway-golang/server/configs/error"
	"api-gateway-golang/server/configs/logger"
	"api-gateway-golang/server/globals"

	"github.com/gin-gonic/gin"
)

type response struct {
	ctx *gin.Context
}

func Ctx(ctx *gin.Context) *response {
	return &response{ctx: ctx}
}

func (response *response) Success(data ...any) {
	if len(response.ctx.Errors) > 0 {
		return
	}
	var result any = data

	if len(data) == 1 {
		result = data[0]
	}
	if len(data) == 0 {
		result = gin.H{}
	}
	responseDataJson, _ := json.MarshalIndent(result, "", "    ")
	logger.TraceLog.Info("[http-response] " + response.ctx.Request.Method + ": " + response.ctx.Request.URL.RequestURI() + " => " + string(responseDataJson))

	response.ctx.JSON(http.StatusOK, result)
	response.ctx.Abort()
}

/**
 * response.Failed()
 * response.Failed("error message") // 第一个参数为错误信息描述，string类型
 * response.Failed("error message", "INNER_SERVER_ERROR") // 第二个参数为错误码，string类型
 * response.Failed("error message", "INNER_SERVER_ERROR", "any data") // 第三个参数为任意数据，可以是任意类型
 * response.Failed("error message", "INNER_SERVER_ERROR", "any data", ["xxx"]) // 第四个参数为错误来源，[]string类型
 */
func (response *response) Failed(data ...any) {
	if len(response.ctx.Errors) > 0 {
		return
	}
	errorException := globals.NewException(data...)
	status := httpError.ErrorCodeMap[errorException.Code]

	responseDataJson, _ := json.MarshalIndent(errorException, "", "    ")
	logger.TraceLog.Info("[http-response] " + response.ctx.Request.Method + ": " + response.ctx.Request.URL.RequestURI() + " => " + string(responseDataJson))
	response.ctx.JSON(status, errorException)
	response.ctx.Abort()
}
