package models

import "gorm.io/gorm"

type Response struct{
	gorm.Model
	Keyword string `gorm:"type:text" json:"content"`
	Message string
}