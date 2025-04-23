package routers

import (
	"api-gateway-golang/server/routers/auth"

	"github.com/gin-gonic/gin"
	ginregex "github.com/jxskiss/ginregex"
)

func RegisterRoutes(r *gin.Engine) {
	var handlers []*ginregex.Matcher

	userManagerMatchers := auth.GetUserManagerAuthMatcher()
	otherMatchers := auth.GetOtherAuthMatcher()
	handlers = append(handlers, userManagerMatchers...)
	handlers = append(handlers, otherMatchers...)

	r.Any("/api/*path", ginregex.Dispatch(handlers...))
}
