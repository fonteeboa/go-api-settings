package services

import (
	types "golang-api-settings/internal/modules/authorization/domain"
)

func (s *AuthorizationService) GetDataService(filter types.Authorization) ([]*types.Authorization, error) {
	return s.repositories.GetData(filter)
}
