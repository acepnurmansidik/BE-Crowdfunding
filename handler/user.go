package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	// tangkap input dari user
	var input user.RegisterUserInput

	// mapping input dari user ke struct RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil{
		c.JSON(http.StatusBadRequest, nil)
	}

	// save input user ke service
	// passing struct sebagai parameter service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil{
		c.JSON(http.StatusBadRequest, nil)
	}

	// melakukan format user
	formatUser := user.FormatUser(newUser, "iniadalahtoken")
	
	// format response API
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatUser)
	// response
	c.JSON(http.StatusOK, response)
}