package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
"cpde": 1001,// 错误码
"msg": xx, // 提示信息
"data": {} //数据
}
*/

/*type ResponseData struct {
	Code ResCode         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResopnseError(c *gin.Context,code ResCode)  {
	rd:= &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}

func ResopnseSuccess(c *gin.Context,data interface{})  {
	rd:= &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK,rd)
}
*/

type ResponseData struct {
	Status string `json:"status"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResopnseError(c *gin.Context,msg interface{})  {
	rd:= &ResponseData{
		Status: "error",
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK,rd)
}

func ResopnseSuccess(c *gin.Context,data interface{})  {

	rd:= &ResponseData{
		Status: "success",
		Msg:  nil,
		Data: data,
	}
	c.JSON(http.StatusOK,rd)
}