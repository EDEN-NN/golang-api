package controllers

import (
	"net/http"

	"gihub.com/EDEN-NN/database"
	"gihub.com/EDEN-NN/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "deu certo caralho"})
}

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		c.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		c.Abort()
		return
	}

	record := database.DB.Create(&user)

	if record.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": record.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if user.ID != 0 {
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Error: User doesn't exists!",
	})
}

func PostUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateFields(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func PatchUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	database.DB.First(&user, id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := models.ValidateFields(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&user).UpdateColumns(user)
	c.JSON(http.StatusAccepted, user)

}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error()})
		return
	}

	database.DB.Delete(&user, id)
	c.JSON(http.StatusAccepted, "Success!")

}
