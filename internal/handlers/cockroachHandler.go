package handlers

import (
	"github.com/gin-gonic/gin"
)

type CockroachHandler interface {
	DetectCockroach(ctx *gin.Context) error
}
