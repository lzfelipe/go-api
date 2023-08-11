package model

import "gorm.io/gorm"

type UserInput struct {
	Username string
	Email string
	Password string
}

type User struct {
	gorm.Model
	Username string
	Email string
	Password string
}

