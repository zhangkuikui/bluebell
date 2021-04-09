package routes

import (
	"bluebell/controllers"
	"github.com/gin-gonic/gin"
)

func UserInit(r *gin.Engine)  {
	v1 := r.Group("")
	{
		v1.POST("/login", controllers.LoginHandler)

	}
}