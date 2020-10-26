package configs

import (
	"cafex/apis"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoute api
func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"GET", "POST", "DELETE"},
	// 	AllowHeaders: []string{"Origin"},
	// }))

	client := r.Group("/v1")
	{
		client.GET("/consumers", apis.GetAllConsumer)
		client.POST("/consumer", apis.GetConsumerByID)
		client.POST("/admin", apis.GetEmployeeByID)
		// client.POST("/user/add", userController.AddUser)
		// client.POST("/user/login", userController.LoginUser)
		// client.POST("/user/update", userController.UpdateUser)
		// client.DELETE("/user/del/:id", userController.DeleteUser)
		// client.POST("/file/test", userController.GetFile)
	}

	return r
}
