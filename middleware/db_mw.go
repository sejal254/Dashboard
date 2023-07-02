package middleware

import (
	"PMD/util"

	"github.com/gin-gonic/gin"
)

func DBConnectionMiddleware(c *gin.Context) {
	const functionName = "middleware.DBConnectionMiddleware"
	util.ConnectToDatabase()
	c.Next()
}
