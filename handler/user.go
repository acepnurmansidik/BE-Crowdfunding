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

	err := c.ShouldBindJSON(&input)
	if err != nil{
		// handler error validation
		errors := helper.FormatValidationError(err)

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
		// error message
		errorMessage := gin.H{"errors": err.Error()}
		// response
		response := helper.APIResponse("Register account failed", http.StatusBadGateway, "failed", errorMessage)
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

func (h *userHandler) Login(c *gin.Context){
	// tampung input dari user
	var input user.LoginUserInput

	// tangkap input dari user lalu masukan ke input
	err := c.ShouldBindJSON(&input)
	// handle error input user
	if err != nil {
		// validation err message
		errors := helper.FormatValidationError(err)
		// masukan error message ke dalam objek errors
		errorMessage := gin.H{"errors": errors}
		// response
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// mapping input dari user ke struct RegisterUserInput
	// save input user ke service
	// passing struct sebagai parameter service
	loginUser, err := h.userService.Login(input)
	// handle error login
	if err != nil{
		// error message
		errorMessage := gin.H{"errors": err.Error()}
		// response
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loginUser,"ini token")

	response := helper.APIResponse("Successfuly loggin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}