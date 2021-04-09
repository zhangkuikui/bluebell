package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}

	// 2. 生成id
	userID := snowflake.GetID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 3. 保存进数据库
	err = mysql.InsertUser(user)
	return

}

func Login(p *models.ParamLogin) (token string,err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是一个指针，就能拿到user.UserID
	if err:=mysql.Login(user);err!=nil{
		return "",err
	}
	// 生成JWT
	return jwt.GenToken(user.UserID,user.Username)

}
