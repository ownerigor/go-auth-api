package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ownerigor/go-api-auth/internal/models"
	"gorm.io/gorm"
)

func MeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		var user models.User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		})
	}
}
