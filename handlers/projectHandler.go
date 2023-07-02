package handlers

import (
	"PMD/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(engine *gin.Engine) {
	r := engine.Group("vasgold")
	r.GET("/project/all", controllers.GetAllProjects)
	r.POST("/project/add", controllers.AddProject)
	r.POST("/project/remove/:project_id", controllers.RemoveProject)
}
