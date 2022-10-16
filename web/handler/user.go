package handler

import (
	"bwastartup/app/user"
	"net/http"
	"strconv"

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
		input.Error = err
		// reload halaman, beserta value yang sudah diisi
		c.HTML(http.StatusOK, "user_new.html", input)
		return
	}

	// mapping
	registerInput := user.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Email = input.Email
	registerInput.Occupation = input.Occupation
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Edit(c *gin.Context){
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	resgiterUser, err := h.userService.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := user.FormUpdateUserInput{}
	input.ID = resgiterUser.ID
	input.Name = resgiterUser.Name
	input.Email = resgiterUser.Email
	input.Occupation = resgiterUser.Occupation

	c.HTML(http.StatusOK, "user_edit.html", input)
}

func (h *userHandler) Update(c *gin.Context){
	// get id di parameter
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	// tangkap value dari form input
	var input user.FormUpdateUserInput
	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		// reload halaman, beserta value yang sudah diisi
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	// binding id nya
	input.ID = id
	// update user
	_, err = h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) NewAvatar(c *gin.Context){
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)


	c.HTML(http.StatusOK, "user_avatar.html", gin.H{"ID": id})
}