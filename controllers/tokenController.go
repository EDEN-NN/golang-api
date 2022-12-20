package controllers

import (
	"net/http"

	"gihub.com/EDEN-NN/auth"
	"gihub.com/EDEN-NN/database"
	"gihub.com/EDEN-NN/models"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	record := database.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": record.Error})
		c.Abort()
		return
	}

	credentialError := user.ComparePassword(request.Password)

	if credentialError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": credentialError.Error()})
		c.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
