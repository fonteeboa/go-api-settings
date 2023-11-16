package domain

import (
	authorizationModels "golang-api-settings/internal/modules/authorization/domain"

	"gorm.io/gorm"
)

// Settings é o modelo correspondente à struct Integration.
type Integration struct {
	gorm.Model
	ApiKey  string
	ApiID   int
	Setting authorizationModels.Authorization
}
