package model

import (
	"gorm.io/gorm"
)

type Alias struct {
	gorm.Model
	Name   string
	Target string `gorm:"index"`
}
