package model

import "github.com/jinzhu/gorm"

type Link struct {
	gorm.Model
	URL       string `json:"address"`
	IsFetched bool   `gorm:"default:false" json:"is_fetched"`
}
