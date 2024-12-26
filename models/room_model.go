package models

import (
	"go_backend/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Room struct {
	Id          string `json:"id"`
	RoomNumber  string `json:"room_number"`
	RoomTypeId  string `json:"room_type_id"`
	IsAvailable string `json:"is_available"`
}

func (b *Room) TableName() string {
	return "rooms"
}

func GetAllRooms(u *[]Room) []Room {
	if err := config.DB.Debug().Find(u).Error; err != nil {
		return *u
	}
	return *u
}

func GetRoom(u *Room, id string) Room {
	u.Id = id
	if err := config.DB.First(u).Error; err != nil {
		return *u
	}

	return *u
}

func AddRoom(user *Room) config.Response {
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

func UpdateRoom(user *Room, id string) config.Response {
	userUpdate := Room{}
	if user.RoomNumber != "" {
		userUpdate.RoomNumber = user.RoomNumber
	}
	if user.RoomTypeId != "" {
		userUpdate.RoomTypeId = user.RoomTypeId
	}
	if user.IsAvailable != "" {
		userUpdate.IsAvailable = user.IsAvailable
	}
	result := config.DB.Debug().Model(&Room{}).Where("id = ?", id).Updates(userUpdate)
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

func DeleteRoom(user *Room, id string) config.Response {
	result := config.DB.Debug().Delete(&Room{}, id)

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
