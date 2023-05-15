package models

import (
	"fmt"

	"github.com/yujen77300/API-test/database"
)

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=8,max=32,pwdvaldation"`
}

type UserVerifiedRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type VerifyResponse struct {
	Success    bool   `json:"success"`
	Reason     string `json:"reason"`
	WrongTimes int    `json:"wrongTimes"`
}

func GetAllUsers() []UserResponse {
	var users []User
	database.DB.Find(&users)
	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, UserResponse{
			Id:       user.Id,
			Username: user.Username,
		})
	}
	return usersResponse
}

func GetUserById(id string) UserResponse {
	var user User
	database.DB.First(&user, id)
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
}

func CreateUser(user UserRequest) UserResponse {
	var existingUser User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return UserResponse{
			Id: 0,
		}
	}

	newUser := User{
		Username: user.Username,
		Password: user.Password,
	}
	database.DB.Create(&newUser)
	return UserResponse{
		Id:       newUser.Id,
		Username: newUser.Username,
	}
}

func VerifyUser(user UserVerifiedRequest, wrongInputTime int) VerifyResponse {
	var existingUser User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return VerifyResponse{
			Success:    false,
			Reason:     fmt.Sprintf("The username '%s' does not exist", user.Username),
			WrongTimes: 0,
		}
	}

	if existingUser.Password != user.Password {
		return VerifyResponse{
			Success:    false,
			Reason:     "Incorrect password",
			WrongTimes: wrongInputTime + 1,
		}
	}

	return VerifyResponse{
		Success:    true,
		Reason:     "",
		WrongTimes: 0,
	}
}


func DeleteUser(id string) bool {
	var user = User{}
	result := database.DB.Where("id = ?", id).Delete(&user)
	return result.RowsAffected != 0
}
