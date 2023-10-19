package services

import (
    "golang-api-settings/internal/domain/apiSystem/types"
    "golang-api-settings/internal/domain/apiSystem/repositories"
)

type ApiSystemService struct {
    repository *repositorys.ApiSystemRepository
}

func NewApiSystemService(repo *repositorys.ApiSystemRepository) *ApiSystemService {
    return &ApiSystemService{repository: repo}
}

func (s *ApiSystemService) GetDataService(filter types.ApiSystem) ([]*types.ApiSystem, error) {
    return s.repository.Get(filter)
}
