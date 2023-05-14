package service

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yujen77300/API-test/models"
)

func GetAllUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.JSON(http.StatusOK, users)

}

func GetUserById(c *gin.Context) {
	user := models.GetUserById(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error 404: User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	user := models.UserRequest{}
	err := c.BindJSON(&user)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "Username") {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  "Please enter a valid username between 3 and 32 characters long.",
			})
			return
		} else if strings.Contains(errMsg, "pwdvaldation"){
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  "Please enter a password with at least 1 uppercase letter, 1 lowercase letter, and 1 number.",
			})
			return
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  "Please enter a password between 8 and 32 characters long.",
			})
			return
		}
	}

	newUser := models.CreateUser(user)
	if newUser.Id == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"reason":  "Username already exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newUser,
	})
}

func DeleteUser(c *gin.Context) {
	user := models.DeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}
