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

// load halaman create user
func (h *userHandler) New(c *gin.Context){
	c.HTML(http.StatusOK, "user_new.html", nil)
} 

// create register user
func (h *userHandler) Create(c *gin.Context){
	var input user.FormCreateUserInput

	// get data
	err := c.ShouldBind(&input)
	if err != nil {
		// 
	}

	// mapping
	registerInput := user.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Email = input.Email
	registerInput.Occupation = input.Occupation
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		// 
	}

	c.Redirect(http.StatusFound, "/users")
}