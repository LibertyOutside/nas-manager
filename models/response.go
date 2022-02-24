package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

type ResContent struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	rc := ResContent{
		Code: 200,
		Msg:  "",
		Data: data,
	}
	r.Ctx.JSON(http.StatusOK, rc)
}

func (r *Result) Error(code int, msg string) {
	rc := ResContent{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	r.Ctx.JSON(http.StatusOK, rc)
}
