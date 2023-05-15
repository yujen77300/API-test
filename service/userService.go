package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/yujen77300/API-test/database"
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
		} else if strings.Contains(errMsg, "pwdvaldation") {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"reason":  "Please enter a password with at least 1 uppercase letter, 1 lowercase letter, and 1 number.",
			})
			return
		} else {
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

func VerifyUser(c *gin.Context) {
	user := models.UserVerifiedRequest{}
	c.BindJSON(&user)
	redisConn := database.RedisPool.Get()
	defer redisConn.Close()

	reply, err := redisConn.Do("GET", user.Username)
	wrongInputTime := 0
	if reply != nil {
		wrongInputTime, _ = redis.Int(reply, err)
	}

	if wrongInputTime == 5 {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"success": false,
			"reason":  "Too many wrong inputs, please try again in 1 minute.",
		})
		return
	}

	verifyResult := models.VerifyUser(user, wrongInputTime)
	if !verifyResult.Success && strings.Contains(verifyResult.Reason, "username") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"reason":  verifyResult.Reason,
		})
		return
	} else if !verifyResult.Success {
		wrongTimeNow := verifyResult.WrongTimes
		if wrongTimeNow == 5 {
			redisConn.Do("SET", user.Username, wrongTimeNow, "EX", "60")
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"reason":  "Too many wrong inputs, please try again in 1 minute.",
			})
			return
		}
		redisConn.Do("SET", user.Username, wrongTimeNow)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  fmt.Sprintf("Invalid password. %d attempts remain before the account is locked out.", 5-wrongTimeNow),
		})
		return
	}
	redisConn.Do("SET", user.Username, verifyResult.WrongTimes)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
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
