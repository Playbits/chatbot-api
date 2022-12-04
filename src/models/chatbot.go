package models

import "gorm.io/gorm"

type Response struct{
	gorm.Model
	Keyword string
	Message string
}