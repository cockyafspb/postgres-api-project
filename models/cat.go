package models

import "gorm.io/gorm"

type Cat struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name     *string `json:"name"`
	Birthday *string `json:"birthday"`
	Color    *string `json:"color"`
	OwnerId  uint    `json:"owner"`
	Owner    Owner   `gorm:"foreignKey:OwnerId;references:ID" json:"-"`
}

func MigrateCats(db *gorm.DB) error {
	err := db.AutoMigrate(&Cat{})
	return err
}
