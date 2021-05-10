package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context)  {

	//ResopnseSuccess(c,"ok")
	//return


	// 1. 获取参数及参数校验
	// c.ShouldBindJSON() //validator --> binding tag
	p:=new(models.Post)
	//if err := c.ShouldBindJSON(p);err!=nil{
	if err:=c.ShouldBindJSON(p);err!=nil{
		//fmt.Println("err:",err)
		zap.L().Error(err.Error())
		ResopnseError(c,"请求参数错误")
		return
	}




	//从c取到当前发请求的用户id
	userID,err:=GetCurrentUserId(c)
	if err !=nil{
		ResopnseError(c,"CodeNeedLogin")
		return
	}
	p.AuthorID=userID

	// 2. 创建帖子
	if err:=logic.CreatePost(p);err!=nil{
		zap.L().Error("logic.CreatePost(p) failed",zap.Error(err))
		ResopnseError(c,"CodeServerBusy")
		return
	}


	// 3. 返回响应
	a:= make([]string, 0)
	ResopnseSuccess(c,a)

}
