package services

import (
	types "golang-api-settings/internal/modules/settings/domain"
)

func (s *SettingsService) GetDataService(filter types.Settings) ([]*types.Settings, error) {
	return s.repositories.GetData(filter)
}
