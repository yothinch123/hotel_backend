package models

import (
	"go_backend/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (b *User) TableName() string {
	return "users"
}

func GetAllUsers(u *[]User) []User {
	if err := config.DB.Debug().Find(u).Error; err != nil {
		return *u
	}
	return *u
}

func GetUser(u *User, id string) User {
	u.Id = id
	if err := config.DB.First(u).Error; err != nil {
		return *u
	}

	return *u
}

func AddUser(user *User) config.Response {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return config.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "CREATED_FAILED",
		}
	}

	return config.Response{
		StatusCode: http.StatusOK,
		Message:    "CREATED_SUCCESS",
	}
}

func UpdateUser(user *User, id string) config.Response {
	userUpdate := User{}
	if user.Name != "" {
		userUpdate.Name = user.Name
	}
	if user.Email != "" {
		userUpdate.Email = user.Email
	}
	if user.Password != "" {
		userUpdate.Password = user.Password
	}
	if user.PhoneNumber != "" {
		userUpdate.PhoneNumber = user.PhoneNumber
	}
	result := config.DB.Debug().Model(&User{}).Where("id = ?", id).Updates(userUpdate)
	if result.Error != nil {
		return config.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "UPDATED_FAILED",
		}
	}

	return config.Response{
		StatusCode: http.StatusOK,
		Message:    "UPDATED_SUCCESS",
	}
}

func DeleteUser(user *User, id string) config.Response {
	result := config.DB.Debug().Delete(&User{}, id)

	if result.Error != nil {
		return config.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "DELETE_FAILED",
			Error:      result.Error,
		}
	}

	return config.Response{
		StatusCode: http.StatusOK,
		Message:    "DELETE_SUCCESS",
	}

}
