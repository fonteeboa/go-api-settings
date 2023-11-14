package services

import (
	types "golang-api-settings/internal/modules/settings/domain"
)

func (s *SettingsService) UpdateData(data *types.Settings) (string, error) {
	// Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
	existingData, err := s.repositories.GetData(types.Settings{Name: data.Name})
	if err != nil {
		return "", err
	}

	if len(existingData) > 0 && existingData[0].ID != data.ID {
		return "common.duplicated", nil
	}

	err = s.repositories.UpdateSettings(data)
	if err != nil {
		return "", err
	}

	return "common.success", nil
}