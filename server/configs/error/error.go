package httperror

const (
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	NOT_FOUND             = "NOT_FOUND"
	URL_NOT_FOUND         = "URL_NOT_FOUND"
	FORBIDDEN             = "FORBIDDEN"
)

/* { [string 错误码]: [int http状态码] } */
var ErrorCodeMap = map[string]int{
	INTERNAL_SERVER_ERROR: 500,
	NOT_FOUND:             404,
	URL_NOT_FOUND:         404,
	FORBIDDEN:             403,
}
