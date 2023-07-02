package handlers

import (
	"PMD/controllers"

	"github.com/gin-gonic/gin"
)


func SetupUserRoutes(engine *gin.Engine) {
	r := engine.Group("/user")
	r.GET("/all", controllers.GetAllUsers)
	r.GET("/:user_id", controllers.GetUserById)
	r.POST("/remove/:user_id", controllers.DeleteUser)
	r.POST("/add", controllers.Adduser)
	r.POST("/login", controllers.Login)
	r.POST("/update/:user_id", controllers.UpdateUserRole)

}

