package routes

import (
	"bluebell/controllers"
	"github.com/gin-gonic/gin"
)

func UserInit(r *gin.Engine,group *gin.RouterGroup)  {
	{
		group.POST("/login", controllers.LoginHandler)
		group.POST("/login", controllers.LoginHandler)
	}
}