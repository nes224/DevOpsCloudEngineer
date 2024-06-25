package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/nginx/middleware"
	"github.com/golang/nginx/services"
)

type Server struct {
	router *gin.Engine
}

var (
	server       *gin.Engine
	ctx          context.Context
	userServices services.ServiceStore
)

func InitServer(router *gin.Engine) *Server {
	services.Store()
	server := &Server{
		router: router,
	}
	server.Start()
	return server
}

func (server *Server) Start() {
	apiGroup := server.router.Group("/api")
	apiGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	})
	apiGroup.POST("/auth/singup", server.SignUpUser)
	apiGroup.POST("/auth/signin", server.SignInUser)
	apiGroup.POST("/auth/logout", middleware.DeserializeUser(), server.LogoutUser)
}
