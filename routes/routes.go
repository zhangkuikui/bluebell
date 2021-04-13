package routes

import (
	"bluebell/controllers"
	"bluebell/middlewares"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode) // 设置成发布模式

	r := gin.New()
	//加载日志中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1:=r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录
	v1.POST("/login", controllers.LoginHandler)

	//可以拆分出去
	//UserInit(r,v1)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用JWT认证中间件

	{
		v1.GET("/community",controllers.CommunityHandler) //
	}

	//r.GET("/ping", func(c *gin.Context) {
	//
	//	controllers.ResopnseSuccess(c,"pong")
	//
	//})


	r.NoRoute(func(c *gin.Context) {
		/*c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})*/
		controllers.ResopnseError(c,"404")
	})
	return r
}
