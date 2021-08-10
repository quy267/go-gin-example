package restaurantmodel

import (
	"go-gin-example/common"
	"strings"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"` // tag
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameCannotBeBlank
	}

	data.Address = strings.TrimSpace(data.Address)

	if data.Address == "" {
		return ErrAddressCannotBeBlank
	}

	return nil
}
