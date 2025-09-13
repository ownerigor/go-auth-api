package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ownerigor/go-api-auth/internal/models"
	"github.com/ownerigor/go-api-auth/pkg/utils"
	"gorm.io/gorm"
)

type SignupInput struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func SignupHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input SignupInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := utils.HashPassword(input.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
			return
		}

		user := models.User{
			Name:     input.Name,
			Username: input.Username,
			Email:    input.Email,
			Password: hashedPassword,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email já registrado"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Usuário criado com sucesso",
		})
	}
}
