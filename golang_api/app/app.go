package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/nginx/config"
	"github.com/golang/nginx/services"
)

type Server struct {
	router *gin.Engine
	store  services.Store
}

var (
	server *gin.Engine
	ctx    context.Context
	store  services.Store
)

func InitServer(router *gin.Engine) *Server {
	server := &Server{
		router: router,
		store:  store,
	}
	server.Start()
	return server
}

func (server *Server) Start() {
	config, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}
	services.StoreDb(config.MONGOURL)
	apiGroup := server.router.Group("/api")
	apiGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	})
	apiGroup.POST("/auth/register", server.SignUpUser)
	
	
}
