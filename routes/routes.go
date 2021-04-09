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

	UserInit(r)



	// 注册
	r.POST("/signup", controllers.SignUpHandler)
	// 登录
	//r.POST("/login", controllers.LoginHandler)





	r.GET("/ping",middlewares.JWTAuthMiddleware(), func(c *gin.Context) {

		controllers.ResopnseSuccess(c,"pong")

	})


	r.NoRoute(func(c *gin.Context) {
		/*c.JSON(http.StatusOK,gin.H{
			"msg":"404",
		})*/
		controllers.ResopnseError(c,"404")
	})
	return r
}
