package domain

import (
	authorizationModels "golang-api-settings/internal/modules/authorization/domain"

	"gorm.io/gorm"
)

// Integration é o modelo correspondente à struct Integration.
type Integration struct {
	gorm.Model
	ApiKey        string
	ApiID         int
	Authorization authorizationModels.Authorization `gorm:"foreignKey:ApiID"`
}
