package util

import "net/http"

func MakeOkResp(data interface{}) (int, interface{}) {
	return http.StatusOK, data
}

type resp struct {
	Code int
	Data interface{}
}

func MakeResp(status int, code int, data interface{}) (int, interface{}) {
	return status, resp{
		code, data,
	}
}
