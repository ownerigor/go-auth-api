package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ownerigor/go-api-auth/internal/config"
	"github.com/ownerigor/go-api-auth/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	db := config.ConnectDataBase(cfg)

	r := gin.Default()

	routes.SetupRoutes(r, db)

	r.Run(":9000")
}
