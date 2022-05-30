package util

import (
	"net/http"
	"strconv"
)

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

func Str2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
