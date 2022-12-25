package controller

import (
	"net/http"

	"github.com/chumnend/pook/internal/entity"
	"github.com/gin-gonic/gin"
)

type userController struct {
	srv entity.UserService
}

// NewController creates a UserController with given UserService
func NewController(srv entity.UserService) entity.UserController {
	return &userController{srv: srv}
}

func (u *userController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is the register endpoint"})
}

func (u *userController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "this is the register endpoint"})
}
