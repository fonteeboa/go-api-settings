package domain

import (
	"gorm.io/gorm"
)

// Settings é o modelo correspondente à struct Settings.
type Settings struct {
	gorm.Model
	Name        string `gorm:"type:text"`
	Description string `gorm:"type:text"`
}
