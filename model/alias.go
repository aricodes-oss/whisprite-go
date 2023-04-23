package model

import (
	"gorm.io/gorm"
)

type CommandAlias struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Target string `gorm:"index"`
}
