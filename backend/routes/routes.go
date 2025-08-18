package routes

import (
	"time"

	"american-empire/backend/api"
	"american-empire/backend/database"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func SetupAPIRoutes(router *gin.Engine) {
	handler := &api.Handler{DB: database.DB}
	store := memory.NewStore()

	loginLimiter := limiter.New(store, limiter.Rate{
		Period: 15 * time.Minute,
		Limit:  5,
	})

	authLimiter := limiter.New(store, limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  30,
	})

	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello from Go!"})
		})
		apiRouter.GET("/", handler.GetEvents)
		apiRouter.POST("/event", handler.GetEvent)
		apiRouter.POST("/contribute", handler.ContributeEvent)
		apiRouter.GET("/tags", handler.GetTags)

		// admin routes
		apiRouter.GET("/auth-me", mgin.NewMiddleware(authLimiter), handler.AuthMe)
		apiRouter.POST("/login", mgin.NewMiddleware(loginLimiter), handler.Login)
		adminRoutes := apiRouter.Group("/")
		adminRoutes.Use(handler.AdminMiddleware())
		{
			adminRoutes.GET("/admin-events", handler.GetAdminEvents)
			adminRoutes.POST("/admin-edit-event", handler.EditEvent)
			adminRoutes.POST("/admin-approve-event", handler.ApproveEvent)
			adminRoutes.POST("/admin-unapprove-event", handler.UnapproveEvent)
			adminRoutes.POST("/admin-upload-photo", handler.UploadPhoto)
			adminRoutes.POST("/admin-upload-youtube", handler.UploadYoutube)
			adminRoutes.POST("/admin-delete-media", handler.DeleteMedia)
			adminRoutes.POST("/admin-delete-source", handler.DeleteSources)
			adminRoutes.POST("/admin-add-source", handler.AddSources)
			adminRoutes.POST("/admin-add-book", handler.AddBook)
		}
	}
}
