package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ownerigor/go-api-auth/internal/handlers"
	"github.com/ownerigor/go-api-auth/internal/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	//Public routes
	r.GET("ping", handlers.PingHandler)
	r.POST("/signup", handlers.SignupHandler(db))
	r.POST("/login", handlers.LoginHashHandler(db))
	r.GET("/login", handlers.LoginJWTHandler(db))

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/me", handlers.MeHandler(db))
		auth.GET("/users", handlers.GetUsersHandler(db))
	}
}
