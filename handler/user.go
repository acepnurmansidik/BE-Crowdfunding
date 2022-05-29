package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	err := c.ShouldBindJSON(&input)
	if err != nil{
		// var utk menampung error
		var errors []string

		// ubah error menjadi error validatior
		for _, e := range err.(validator.ValidationErrors){
			// simpan setiap error string validator ke dalam slice errors
			errors = append(errors, e.Error())
		}

		// masukan var errors ke dalam object
		errorMessage := gin.H{"errors": errors}
		
		// response
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	// mapping input dari user ke struct RegisterUserInput
	// save input user ke service
	// passing struct sebagai parameter service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil{
		response := helper.APIResponse("Register account failed", http.StatusBadGateway, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// melakukan format user
	formatUser := user.FormatUser(newUser, "iniadalahtoken")

	// format response API
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatUser)
	// response
	c.JSON(http.StatusOK, response)
}