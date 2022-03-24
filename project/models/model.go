package models

import "gorm.io/gorm"

type (
	Employee struct {
		gorm.Model
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Age      int64  `json:"age"`
		Location string `json:"address"`
	}
)
