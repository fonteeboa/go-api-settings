package services

import (
	types "golang-api-settings/internal/modules/integration/domain"
)

func (s *IntegrationService) UpdateData(data *types.Integration) (string, error) {
	// Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
	existingData, err := s.repositories.GetData(types.Integration{ApiID: data.ApiID})
	if err != nil {
		return "", err
	}

	if len(existingData) > 0 && existingData[0].ID != data.ID {
		return "common.duplicated", nil
	}

	err = s.repositories.Update(data)
	if err != nil {
		return "", err
	}

	return "common.success", nil
}
