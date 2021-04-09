package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)
const CtxUserIdKey = "userID"
const CtxUserNameKey = "username"

func GetCurrentUserId(c *gin.Context) (userID int64,err error) {
	uid,ok:=c.Get(CtxUserIdKey)
	if !ok{
		err = errors.New("用户未登录")
		return
	}
	userID,ok=uid.(int64)
	if !ok{
		err = errors.New("用户未登录")
		return
	}
	return
}

func GetCurrentUserName(c *gin.Context) (username string,err error) {
	name,ok:=c.Get(CtxUserNameKey)
	if !ok{
		err = errors.New("用户未登录")
		return "",err
	}
	username,ok=name.(string)
	if !ok{
		err = errors.New("用户未登录")
		return "",err
	}
	return username,err
}
