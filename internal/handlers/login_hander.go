package handlers

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ownerigor/go-api-auth/internal/models"
	"github.com/ownerigor/go-api-auth/internal/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHashHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		var user models.User
		if err := db.Where("username = ?", body.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		hash := sha1.New()
		hash.Write([]byte(body.Username + body.Password + strconv.FormatInt(time.Now().Unix(), 10)))
		tokenHash := hex.EncodeToString(hash.Sum(nil))

		db.Create(&models.UserToken{
			UserID:    user.ID,
			TokenHash: tokenHash,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		})

		c.JSON(http.StatusOK, gin.H{"bearer": tokenHash})
	}
}

func LoginJWTHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		var token models.UserToken
		if err := db.Where("token_hash = ? AND expires_at > ?", authHeader, time.Now()).First(&token).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired hash token"})
			return
		}

		jwtToken, err := services.GenerateJWT(token.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"bearer": jwtToken})
	}
}
