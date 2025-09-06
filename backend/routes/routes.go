package routes

import (
	"american-empire/backend/api"
	"american-empire/backend/database"

	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router *gin.Engine) {
	handler := &api.Handler{DB: database.DB}

	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello from Go!"})
		})
		apiRouter.GET("/", handler.GetEvents)
		apiRouter.POST("/event", handler.GetEvent)
		apiRouter.POST("/contribute", handler.ContributeEvent)
		apiRouter.GET("/tags", handler.GetTags)
		apiRouter.GET("/book/:id", handler.GetBook)
	}
}
