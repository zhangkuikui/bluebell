package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
