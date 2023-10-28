package book

import (
	"gorm.io/gorm"
)

type Books struct {
	ID 		uint `gorm:"primary key;autoIncrement" json:"id"`
	Title 	*string `json:"title"`
	Author 	*string `json:"author"`
	Rating 	*int `json:"rating"`
}

func Migration(db *gorm.DB)error{
	err := db.AutoMigrate(&Books{})
	return err
}

