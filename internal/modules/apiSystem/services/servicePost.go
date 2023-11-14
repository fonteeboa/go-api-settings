package services

import (
	"fmt"
	types "golang-api-settings/internal/modules/apiSystem/domain"
)

func (s *ApiSystemService) InsertData(data *types.ApiSystem) (string, error) {
	// Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
	existingData, err := s.repositories.GetData(types.ApiSystem{ApiID: data.ApiID})
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
