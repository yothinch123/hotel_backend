package models

import (
	"fmt"
	"go_backend/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "user"
const dbpass = "password"
const dbname = "hotel_booking"

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

func UpdateUser(user *User, id string) Response {
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
	result := config.DB.Debug().Delete(&User{}, id)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	if result.Error != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "DELETE_FAILED",
			Error:      result.Error,
		}
	}

	return Response{
		StatusCode: http.StatusOK,
		Message:    "DELETE_SUCCESS",
	}

}
