package router

import (
	"github.com/gin-gonic/gin"
	//_ "PoliticianRating/docs"
	"user-service/pkg/controllers"
)

func InitRouter(r *gin.Engine) {
	ratingRoutes := r.Group("/users")
	{
		ratingRoutes.GET("/", controllers.ReadAllUsers)
		ratingRoutes.GET("/:id", controllers.ReadOneUser)
		ratingRoutes.DELETE("/:id", controllers.DeleteUser)
		ratingRoutes.POST("/", controllers.CreateUser)
		ratingRoutes.PUT("/:id", controllers.UpdateUser)
		//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}
