package domain

import (
	"gorm.io/gorm"
)

// Settings é o modelo correspondente à struct ApiSystem.
type ApiSystem struct {
    gorm.Model
    Name   string `gorm:"type:text"`
}