package handler

import (
	"bwastartup/app/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler{
	return &sessionHandler{userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginUserInput

	// binding data dari user
	err := c.ShouldBind(&input)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// lakukan pengecekan ketika akan login
	user, err := h.userService.Login(input)
	// cek jika ada error atau rolenya bukan admin
	if err != nil || user.Role != "admin" {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	
	// simpan data user ke session
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	// redirect ke halaman user jika berhail login
	c.Redirect(http.StatusFound, "/users")
}