package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Ping(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	CreateGame(ctx *gin.Context)
	GetGame(ctx *gin.Context)
	GetBoardGame(ctx *gin.Context)
	Turn(ctx *gin.Context)
}
