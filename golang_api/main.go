package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang/nginx/app"
	"github.com/golang/nginx/config"
	"github.com/golang/nginx/services"
)

var (
	server  *gin.Engine
	service services.ServiceStore
	ctx     context.Context
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server = gin.Default()
	app.InitServer(server)
	server.Run(config.Port)
}
