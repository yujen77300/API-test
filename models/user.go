package models

import (
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

type UserResponse struct {
	Id			 uint   `json:"id"`
	Username string `json:"username"`
}

func  GetAllUsers()[]UserResponse{
	var users []User
	database.DB.Find(&users)
	var usersResponse []UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, UserResponse{
			Id: user.Id,
			Username: user.Username,
		})
	}
	return usersResponse
}

func GetUserById(id string)UserResponse{
	var user User
	database.DB.First(&user, id)
	return UserResponse{
		Id: user.Id,
		Username: user.Username,
	}
}

func CreateUser(user UserRequest)UserResponse{
		var existingUser User
    if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
        return UserResponse{
            Id : 0,
        }
    }

	newUser := User{
		Username: user.Username,
		Password: user.Password,
	}
	database.DB.Create(&newUser)
	return UserResponse{
		Id: newUser.Id,
		Username: newUser.Username,
	}
}

// func UpdateUser(id string, user UserRequest)UserResponse{
// 	database.DB.Model(&user).Where("id = ?", id).Updates(user)
// 	userId,_ := strconv.Atoi(id)
// 	return UserResponse{
// 		Id: uint(userId),
// 		Username: user.Username,
// 	}
// }

func DeleteUser(id string)bool{
	var user = User{}
	result := database.DB.Where("id = ?",id ).Delete(&user)
	return result.RowsAffected != 0
}