// response
package models

type Response struct {
	Success int         `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 出错返回
func GenerateError(msg string) Response {
	errorMsg := Response{1, msg, nil}
	return errorMsg
}

// 正确返回
func GenerateSuccess(msg string, data interface{}) Response {
	errorMsg := Response{0, msg, data}
	return errorMsg
}
