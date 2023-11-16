package services

import (
	types "golang-api-settings/internal/modules/authorization/domain"
)

func (s *AuthorizationService) InsertData(data *types.Authorization) (string, error) {
	// Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
	existingData, err := s.repositories.GetData(types.Authorization{Name: data.Name})
	if err != nil {
		return "", err
	}

	if len(existingData) > 0 {
		return "common.duplicated", nil
	}

	err = s.repositories.CreateAuthorization(data)
	if err != nil {
		return "", err
	}

	return "common.success", nil
}
