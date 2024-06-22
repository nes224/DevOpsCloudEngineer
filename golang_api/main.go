package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/nginx/config"
)

func main() {
	_, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := gin.Default()
	router := server.Group("/api")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Hello Golang Api."})
	})

	server.Run(":" + "8000")
}
