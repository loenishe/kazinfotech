package handler

import (
	"kazinfotech/internal/model"
)

func MakeResponse(data interface{}, statusCode int, statusMessage string) *model.Response {
	var resp model.Response
	resp.Data = data
	resp.ResponseCode = statusCode
	resp.ResponseMessage = statusMessage
	return &resp
}
