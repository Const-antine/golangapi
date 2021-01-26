package mappings

import (
	"goapi/controllers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/users/:id", controllers.GetUserDetail)
		v1.GET("/users/", controllers.GetUser)
		v1.POST("/users", controllers.PostUser)
	}
	// Swagger interface
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
