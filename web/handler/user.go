package handler

import (
	"bwastartup/app/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUerHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// function untuk load file html
func (h *userHandler) Index(c *gin.Context){
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return                
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}