package domain

import (
	"gorm.io/gorm"
)

// Authorization é o modelo correspondente à struct Authorization.
type Authorization struct {
	gorm.Model
	Name        string `gorm:"type:text"`
	Description string `gorm:"type:text"`
}
