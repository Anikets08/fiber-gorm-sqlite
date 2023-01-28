package models

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name       string
	Email      string `gorm:"uniqueIndex"`
	Phone      string
	Categories []Category
}

type Category struct {
	gorm.Model
	Name         string
	MenuItems    []MenuItem
	RestaurantID uint
}

type MenuItem struct {
	gorm.Model
	Name       string
	Price      float32
	CategoryID uint
}
