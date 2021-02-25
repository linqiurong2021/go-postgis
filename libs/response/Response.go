package response

import (
	"encoding/json"
)

// Response Response
type Response struct {
	Code int64
	Msg  string
	Data interface{}
}

// New New
func New(code int64, msg string, data interface{}) (result interface{}, err error) {
	response := &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//
	bytes, err := json.Marshal(response)
	return string(bytes), err
}

// Success Success
func Success(msg string, data interface{}) (result interface{}, err error) {
	return New(200, msg, data)
}

// Error Error
func Error(msg string, data interface{}) (result interface{}, err error) {

	return New(400, msg, data)
}
