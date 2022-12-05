package controllers

import (
	"jwt_example/config"
	"jwt_example/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// If the request body is not valid, gin will return a 400 error
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error hashing password",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating user",
		})
		return
	}

	// Return the user
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// If the request body is not valid, gin will return a 400 error
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// Find the user
	var user models.User
	result := config.DB.First(&user, "email = ?", body.Email)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error signing token",
		})
		return
	}

	// Return the token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*3600, "", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if exists {
		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
		})
	}

}

func Logout(c *gin.Context) {
	c.SetCookie("user", "", -1, "", "localhost", false, true)
	c.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "User is logged out",
	})
}
