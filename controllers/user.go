package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context)  {
	// 1. 获取参数和参数校验
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err:=c.ShouldBindJSON(p);err!=nil{
		//请求参数有误
		zap.L().Error("SignUp with invalid param",zap.Error(err))
		/*c.JSON(http.StatusOK,gin.H{
			//"msg":"请求参数有误",
			"msg":err.Error(),
		})*/
		ResopnseError(c,"请求参数有误")
		return
	}

	/*//手动对请求参数进行详细的业务规则校验
	if len(p.Username)==0||len(p.Password)==0||len(p.RePassword)==0||p.RePassword!=p.Password{
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusOK,gin.H{
			"msg":"请求参数有误",
		})
		return
	}*/

	// 2. 业务处理
	// 传结构体最好传指针提高性能
	//logic.SignUp(&p)
	if err:=logic.SignUp(p);err!=nil{
		/*c.JSON(http.StatusOK,gin.H{
			//"msg":err.Error(),
			"msg":"创建用户失败",
		})
		return*/
		ResopnseError(c,"创建用户失败")
	}

	// 3. 返回响应
	/*c.JSON(http.StatusOK,gin.H{
		"msg":"success",
	})*/
	ResopnseSuccess(c,nil)
}

func LoginHandler(c *gin.Context)  {
	// 1. 获取参数和参数校验
	p := new(models.ParamLogin)
	if err:=c.ShouldBindJSON(p);err!=nil{
		zap.L().Error("Login with invalid param",zap.Error(err))
		/*c.JSON(http.StatusOK,gin.H{
			//"msg":"请求参数有误",
			"msg":err.Error(),
		})*/
		ResopnseError(c,err.Error())
		return
	}

	// 2. 业务处理
	token,err:=logic.Login(p)
	if err!=nil{
		zap.L().Error("Login failed",zap.String("username",p.Username),zap.Error(err))
		/*c.JSON(http.StatusOK,gin.H{
			"msg":"用户名或密码不正确",
		})*/
		ResopnseError(c,"用户名或密码不正确")
		return
	}

	// 3. 返回响应
	/*c.JSON(http.StatusOK,gin.H{
		"msg":"success",
	})*/


	ResopnseSuccess(c,token)

}