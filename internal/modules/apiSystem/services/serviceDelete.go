package services

import (
    types "golang-api-settings/internal/modules/apiSystem/domain"
)

func (s *ApiSystemService) Delete(filter types.ApiSystem) ([]*types.ApiSystem, error) {
    return s.repositories.GetData(filter)
}