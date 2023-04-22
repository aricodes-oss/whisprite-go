package model

import (
	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	Content string
	AddedBy string
}
