package restaurantmodel

import (
	"errors"
	"go-gin-example/common"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"` // tag
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrNameCannotBeBlank    = errors.New("name cannot be blank")
	ErrAddressCannotBeBlank = errors.New("address cannot be blank")
)
