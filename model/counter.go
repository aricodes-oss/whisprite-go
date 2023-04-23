package model

import (
	"gorm.io/gorm"
)

type Counter struct {
	gorm.Model

	Name          string `gorm:"unique"`
	Value         uint64
	Contributions []CounterContribution
}

type CounterContribution struct {
	gorm.Model

	CounterID uint
	UserID    uint
}
