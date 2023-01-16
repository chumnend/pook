package controller

import (
	"net/http"

	"github.com/chumnend/pook/internal/entity"
	"github.com/gin-gonic/gin"
)

type userController struct {
	srv entity.UserService
}

// NewUserController creates a UserController with given UserService
func NewUserController(srv entity.UserService) entity.UserController {
	return &userController{srv: srv}
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *userController) Register(c *gin.Context) {
	// validate input
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create new user
	user := entity.User{}
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password

	// save user
	if err := u.srv.Save(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *userController) Login(c *gin.Context) {
	// validate input
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user exists
	user, err := u.srv.FindByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username and/or password"})
		return
	}

	// check password
	if err := u.srv.ComparePassword(user, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username and/or password"})
		return
	}

	// generate token
	token, err := u.srv.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong, please try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
