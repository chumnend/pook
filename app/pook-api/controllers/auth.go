package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint"})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "this is the login endpoint"})
}
