package gorm

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name   string
	Gender int
	Phone  string
}
