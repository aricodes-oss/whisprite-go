package model

import (
	"gorm.io/gorm"
)

type CommandAlias struct {
	gorm.Model
	Name   string
	Target string `gorm:"index"`
}
