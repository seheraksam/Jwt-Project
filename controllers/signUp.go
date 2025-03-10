package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Jwt-Project/initializers"
	"github.com/seheraksam/Jwt-Project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var body struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to load body"})
	}
	if !strings.HasSuffix(body.Email, ".com") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email must end with .com"})
		return
	}
	var existingUser models.User
	err := initializers.Client.Database("jwt-project").Collection("users").FindOne(c.Request.Context(), bson.M{"email": body.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Name: "seher", Email: body.Email, Password: string(hash), CreatedAt: time.Now(), UpdatedAt: time.Now(), ID: primitive.NewObjectID()}
	result, err := initializers.Client.Database("jwt-project").Collection("users").InsertOne(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user to mongo"})
		fmt.Println(err)
		return
	}
	fmt.Println(result.InsertedID)
	c.JSON(http.StatusOK, gin.H{"Succes": "save succesfull save to mongo"})
}
