package configs

type gatewayConfig struct {
	WhitePrefixes []string `yaml:"white_prefixes"`
}

var config = gatewayConfig{
	WhitePrefixes: []string{
		"/api/v1/",
		"/api/usermanager/v1/",
		"/api/usermanageradm/v1/",
	},
}

func GetGatewayConfig() gatewayConfig {
	return config
}
