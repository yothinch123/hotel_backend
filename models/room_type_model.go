package models

import (
	"go_backend/config"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type RoomType struct {
	Id            string  `json:"id"`
	TypeName      string  `json:"type_name"`
	Description   string  `json:"description"`
	PricePerNight float32 `json:"price_per_night"`
}

func (b *RoomType) TableName() string {
	return "room_types"
}

func GetAllRoomTypes(u *[]RoomType) []RoomType {
	if err := config.DB.Debug().Find(u).Error; err != nil {
		return *u
	}
	return *u
}

func GetRoomType(u *RoomType, id string) RoomType {
	u.Id = id
	if err := config.DB.First(u).Error; err != nil {
		return *u
	}

	return *u
}

func AddRoomType(user *RoomType) config.Response {
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

func UpdateRoomType(user *RoomType, id string) config.Response {
	userUpdate := RoomType{}
	if user.TypeName != "" {
		userUpdate.TypeName = user.TypeName
	}
	if user.Description != "" {
		userUpdate.Description = user.Description
	}
	if user.PricePerNight != 0 {
		userUpdate.PricePerNight = user.PricePerNight
	}
	result := config.DB.Debug().Model(&RoomType{}).Where("id = ?", id).Updates(userUpdate)
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

func DeleteRoomType(user *RoomType, id string) config.Response {
	result := config.DB.Debug().Delete(&RoomType{}, id)

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
