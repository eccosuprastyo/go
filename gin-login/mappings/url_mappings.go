package mappings

import (
	"goapi/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/users/:id", controllers.GetUserDetail)
		v1.GET("/users/", controllers.GetUser)
		v1.POST("/login/", controllers.Login)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.POST("/users", controllers.PostUser)
	}
}
