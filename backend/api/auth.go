package api

import (
	"fmt"
	"good-guys/backend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) AuthMe(c *gin.Context) {
	tokenString, err := c.Cookie("good_guys_auth_token")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authenticated": false})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			jwtSecret = "fallback-secret-key"
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authenticated": false})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusOK, gin.H{"authenticated": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authenticated": true})
}

func (h *Handler) Login(c *gin.Context) {
	var request models.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	envUsername := os.Getenv("ADMIN_USERNAME")
	envPassword := os.Getenv("ADMIN_PASSWORD_HASHED")

	if request.Username != envUsername || bcrypt.CompareHashAndPassword([]byte(envPassword), []byte(request.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateJWT()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}

	isProduction := os.Getenv("GIN_MODE") == "release"

	maxAge := 90 * 24 * 60 * 60
	c.SetCookie(
		"good_guys_auth_token",
		token,
		maxAge,
		"/",
		"",
		isProduction,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Authentication successful",
	})
}

func generateJWT() (string, error) {
	claims := jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour * 24 * 90).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "fallback-secret-key" // Only use in development
	}
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}
