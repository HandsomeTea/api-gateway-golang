package auth

import (
	"api-gateway-golang/server/services"

	ginregex "github.com/jxskiss/ginregex"
)

func GetUserManagerAuthMatcher() []*ginregex.Matcher {
	methods := []string{"GET", "POST", "PUT", "DELETE"}
	matchers := make([]*ginregex.Matcher, 0, len(methods)*2)

	for _, method := range methods {
		matchers = append(matchers, ginregex.NewMatcher(method, `^/api/usermanager/v(\d+)/.*`, services.UserManagerAuthService.LoginAuth))
		matchers = append(matchers, ginregex.NewMatcher(method, `^/api/usermanageradm/v(\d+)/.*`, services.UserManagerAuthService.AdminAuth))
	}

	return matchers
}
