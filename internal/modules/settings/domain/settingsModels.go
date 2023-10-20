package domain

import (
	"gorm.io/gorm"
)

// Settings é o modelo correspondente à struct Settings.
type Settings struct {
    gorm.Model
    ApiKey    string `gorm:"type:text"`
    ApiId     uint   `gorm:"TYPE:integer REFERENCES api_systems"`
}