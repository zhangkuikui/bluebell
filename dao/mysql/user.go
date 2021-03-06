package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "zkk.com"

// CheckUserExist 检查指定用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return  err
	}
	if count > 0{
		return errors.New("用户已存在")
	}
	return

}
// Login 用户登录
func Login(user *models.User)(err error)  {
	oPassword:=user.Password //用户登录的密码
	sqlStr := `select user_id,username,password from user where username=?`

	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows{
		return errors.New("用户不存在")
	}
	if err != nil {
		//查询数据库失败
		return  err
	}
	//判断密码是否正确
	if encryptPassword(oPassword)!=user.Password{
		return errors.New("密码不正确")
	}
	return
}

// InsertUser 像数据库中写入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password=encryptPassword(user.Password)
	//sql语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password);
	return
}

func encryptPassword(oPassword string) string  {
	h:=md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
