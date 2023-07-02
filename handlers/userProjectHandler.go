package handlers

import (
	"PMD/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserprojectRoutes(engine *gin.Engine) {

	p := engine.Group("/userp_roject")
	p.GET("/userp/:user_id", controllers.GetProjectsOfUserID)

	r := engine.Group("user_project")
	r.POST("/project", controllers.AssignProjectByID)

}
