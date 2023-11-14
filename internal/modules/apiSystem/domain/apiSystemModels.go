package domain

import (
	settingsModels "golang-api-settings/internal/modules/settings/domain"

	"gorm.io/gorm"
)

// Settings é o modelo correspondente à struct ApiSystem.
type ApiSystem struct {
	gorm.Model
	ApiKey  string
	ApiID   int
	Setting settingsModels.Settings
}
