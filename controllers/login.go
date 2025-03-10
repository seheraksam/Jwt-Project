package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/seheraksam/Jwt-Project/initializers"
	"github.com/seheraksam/Jwt-Project/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Request body parse et
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to load body"})
		return
	}

	// Environment değişkenlerini yükle
	godotenv.Load()
	secretKey := os.Getenv("SECRET_KEY")

	// Kullanıcıyı MongoDB'den çek
	var loggedUser models.User
	err := initializers.Client.Database("jwt-project").Collection("users").FindOne(c.Request.Context(), bson.M{"email": body.Email}).Decode(&loggedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Hashlenmiş şifreyi karşılaştır
	err = bcrypt.CompareHashAndPassword([]byte(loggedUser.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// JWT Token oluştur
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loggedUser.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 gün geçerli
	})

	// Tokeni secret key ile imzala
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// Başarılı giriş yanıtı
	c.JSON(http.StatusOK, gin.H{})
}
