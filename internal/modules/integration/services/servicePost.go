package services

import (
	"fmt"
	types "golang-api-settings/internal/modules/integration/domain"
)

func (s *IntegrationService) InsertData(data *types.Integration) (string, error) {
	// Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
	existingData, err := s.repositories.GetData(types.Integration{ApiID: data.ApiID})
	if err != nil {
		return "", err
	}

	fmt.Println(existingData)

	if len(existingData) > 0 {
		return "common.duplicated", nil
	}

	err = s.repositories.Create(data)
	if err != nil {
		return "", err
	}

	return "common.success", nil
}
