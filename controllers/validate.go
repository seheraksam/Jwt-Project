package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "I'm logged in."})
}
