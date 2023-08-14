package models

import "gorm.io/gorm"

type Owner struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name     *string `json:"name"`
	Birthday *string `json:"birthday"`
}

func MigrateOwners(db *gorm.DB) error {
	err := db.AutoMigrate(&Owner{})
	return err
}
