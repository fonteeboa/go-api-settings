package services

import (
	types "golang-api-settings/internal/modules/integration/domain"
)

func (s *IntegrationService) Delete(filter types.Integration) error {
	return s.repositories.DeleteByID(filter.ID)
}
