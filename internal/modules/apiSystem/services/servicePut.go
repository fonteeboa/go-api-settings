package services

import (
	types "golang-api-settings/internal/modules/apiSystem/domain"
)

func (s *ApiSystemService) UpdateData(data *types.ApiSystem) (string, error) {
	// Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
	existingData, err := s.repositories.GetData(types.ApiSystem{ApiID: data.ApiID})
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
