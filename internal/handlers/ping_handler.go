package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	})
}
