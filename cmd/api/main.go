package main

import (
	"github.com/ownerigor/go-api-auth/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	config.ConnectDataBase(cfg)
}
