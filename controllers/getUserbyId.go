package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/seheraksam/Jwt-Project/initializers"
	"github.com/seheraksam/Jwt-Project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByID(c *gin.Context, userID string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Kullanıcıyı MongoDB'den çek
	var user models.User
	err = initializers.Client.Database("jwt-project").Collection("users").FindOne(
		c.Request.Context(),
		bson.M{"_id": objID},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
