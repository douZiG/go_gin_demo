package routers

import (
	"../controllers"
	"./Cors"
	"github.com/gin-gonic/gin"
)

func ConfigRouters() *gin.Engine {
	//router := gin.Default()
	GRouter := gin.Default()

	GRouter.Use(Cors.CORS(Cors.Options{Origin: "http://localhost:3000"}))

	GRouter.Use(gin.Recovery())
	user := GRouter.Group("/user")
	{
		user.GET("queryUser/*name", controllers.QueryUser)
		user.POST("addUser", controllers.AddUser)
		user.PUT("updateUser", controllers.UpdateUser)
		user.DELETE("deleteUser", controllers.DeleteUser)

	}

	return GRouter
}
