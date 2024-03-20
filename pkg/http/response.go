package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	c *gin.Context
}

type Response interface {
	ResposeCreatedWithId(id string)
	ResposeAbortWithMessage(err error, msg string)
	ResposeBadRequest(msg string)
	ResposeOKWithJSON(obj any)
}

func NewResponse(c *gin.Context) Response {
	return &response{
		c: c,
	}
}

func (r *response) ResposeCreatedWithId(id string) {
	r.c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (r *response) ResposeAbortWithMessage(err error, msg string) {
	r.c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": msg})
}

func (r *response) ResposeBadRequest(msg string) {
	r.c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": msg})
}

func (r *response) ResposeOKWithJSON(obj any) {
	r.c.JSON(http.StatusOK, obj)
}
