package handlers

import (
	"PMD/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	engine.Use(middleware.DBConnectionMiddleware)
	engine.Use(cors.Default())
	SetupUserRoutes(engine)
	SetupProjectRoutes(engine)
	SetupUserprojectRoutes(engine)
}
