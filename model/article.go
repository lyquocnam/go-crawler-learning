package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model

	URL     string `gorm:"unique;not null" json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
