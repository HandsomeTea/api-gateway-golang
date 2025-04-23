package env

import "os"

var defaultEnv = map[string]string{
	"GO_ENV":          "development",
	"PORT":            "8050",
	"LOG_LEVEL":       "debug", // log level value: debug info warn error dpanic panic fatal
	"TRACE_LOG_LEVEL": "debug",
}

func GetEnv(key string) string {
	result := os.Getenv(key)

	if result != "" {
		return result
	}
	return defaultEnv[key]
}
