package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func NewUerHandler() *userHandler {
	return &userHandler{}
}

// function untuk load file html
func (h *userHandler) Index(c *gin.Context){
	c.HTML(http.StatusOK, "user_index.html", nil)
}