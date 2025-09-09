package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ownerigor/go-api-auth/internal/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/signup", handlers.SignupHandler(db))
}
