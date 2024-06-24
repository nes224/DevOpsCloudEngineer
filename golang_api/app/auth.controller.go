package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/nginx/models"
	"github.com/golang/nginx/services"
)

func (server *Server) SignUpUser(ctx *gin.Context) {
	var user *models.SignUpInput
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if user.Password != user.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}
	newUser, err := services.SignUpUser(user,ctx)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": models.FilteredResponse(newUser)})
}
