package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post)(err error)  {
	// 1. 生成postid
	p.ID=snowflake.GetID()


	// 2. 保存到数据库
	return mysql.CreatePost(p)

	// 3. 返回
}
