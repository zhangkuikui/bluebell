package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// ---跟社区相关
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区 (community_id,community_name)
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResopnseError(c, err.Error()) //数据库查询出错
		return
	}
	ResopnseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context)  {
	// 1. 获取社区id
	id_str:=c.Param("id")
	id,err:=strconv.ParseInt(id_str,10,64)
	if err!=nil{
		ResopnseError(c,err.Error())
		return
	}
	// 2. 根据id获取详情
	data,err:=logic.GetCommunityDetail(id)
	if err!=nil{
		zap.L().Error("logic.GetCommunityDetail failed",zap.Error(err))
		ResopnseError(c,err.Error())
		return
	}
	ResopnseSuccess(c,data)

}
