package services

import (
	"golang-api-settings/internal/modules/integration/repositories"
)

type IntegrationService struct {
	repositories *repositories.IntegrationRepository
}

// NewIntegrationService creates a new instance of the IntegrationService struct.
// It takes a pointer to an IntegrationRepository as a parameter and returns a pointer to the IntegrationService.
func NewIntegrationService(repo *repositories.IntegrationRepository) *IntegrationService {
	return &IntegrationService{repositories: repo}
}
