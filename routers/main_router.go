package routers

import (
	"go_backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")

	{
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users/:id", controllers.GetUserByID)
		v1.POST("/users", controllers.CreateUser)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)

		v1.GET("/room", controllers.GetRoom)
		v1.GET("/room/:id", controllers.GetRoomByID)
		v1.POST("/room", controllers.CreateRoom)
		v1.PUT("/room/:id", controllers.UpdateRoom)
		v1.DELETE("/room/:id", controllers.DeleteRoom)
	}

	return r
}
