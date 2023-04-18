package model

import (
	"gorm.io/gorm"
)

type Alias struct {
	gorm.Model
	Name   string
	Target string `gorm:"index"`
}

type Quote struct {
	gorm.Model
	Content string
	AddedBy string
}
