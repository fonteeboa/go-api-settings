package services

import (
    types "golang-api-settings/internal/modules/apiSystem/domain"
)

func (s *ApiSystemService) InsertData(data *types.ApiSystem) (string, error) {
    // Verifique se já existe um registro com as mesmas informações (verifique duplicatas)
    existingData, err := s.repositories.GetData(types.ApiSystem{Name: data.Name})
    if err != nil {
        return "", err
    }

    if len(existingData) > 0 {
        return "common.duplicated", nil
    }

    if data.ID == 0 {
        err = s.repositories.Create(data);
    } else {
        err = s.repositories.Update(data);
    }

    if err != nil {
        return "", err
    }

    return "common.success", nil
}
