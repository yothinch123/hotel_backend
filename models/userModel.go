package models

import (
	"go_backend/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "user"
const dbpass = "password"
const dbname = "hotel_booking"

func GetAllUsers(u *[]User) []User {
	if err := config.DB.Find(u).Error; err != nil {
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

func AddUser(user *User) Response {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "CREATED_FAILED",
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Message:    "CREATED_SUCCESS",
	}
}

func UpdateUser(user *User) Response {
	result := config.DB.Model(&User{}).Where("id = ?", user.Id).Updates(user)

	if result.Error != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "UPDATED_FAILED",
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Message:    "UPDATED_SUCCESS",
	}
}

func DeleteUser(user *User, id string) Response {
	result := config.DB.Delete(&user)

	if result.Error != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "DELETE_FAILED",
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Message:    "DELETE_SUCCESS",
	}

}
