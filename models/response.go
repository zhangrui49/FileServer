// response
package models

type Response struct {
	Success int         `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func init() {

}

func GenerateError(msg string) Response {
	errorMsg := Response{1, msg, nil}
	return errorMsg
}

func GenerateSuccess(msg string, data interface{}) Response {
	errorMsg := Response{0, msg, data}
	return errorMsg
}
