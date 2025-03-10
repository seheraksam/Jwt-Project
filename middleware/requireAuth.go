package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/seheraksam/Jwt-Project/controllers"
)

// JWT doğrulama ve yetkilendirme middleware
func RequireAuth(c *gin.Context) {
	// Cookie'den JWT token'ı al
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie not found"})
		return
	}

	// JWT token'ı çözümle
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// HMAC dışında bir imzalama metodu kullanılmışsa hata ver
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Secret key'i döndür
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// Token doğrulama başarısız olduysa yetkisiz erişimi engelle
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Token içindeki claims'leri al
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	// Token süresi dolmuş mu kontrol et
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		return
	}

	// Kullanıcı ID'sini token'dan al
	userID, ok := claims["sub"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	// Kullanıcıyı MongoDB'den al
	user, err := controllers.GetUserByID(c, userID)
	if err != nil || user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Kullanıcıyı context'e ekle (ilerleyen middleware'ler için)
	c.Set("user", user)

	// Middleware işlemine devam et
	c.Next()
}
