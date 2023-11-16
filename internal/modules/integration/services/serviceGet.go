package services

import (
	types "golang-api-settings/internal/modules/integration/domain"
)

func (s *IntegrationService) GetDataService(filter types.Integration) ([]*types.Integration, error) {
	return s.repositories.GetData(filter)
}
