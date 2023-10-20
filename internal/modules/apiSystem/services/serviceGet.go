package services

import (
    "golang-api-settings/internal/modules/apiSystem/types"
    "golang-api-settings/internal/modules/apiSystem/repositories"
)

type ApiSystemService struct {
    repository *repositorys.ApiSystemRepository
}

// NewApiSystemService creates a new instance of the ApiSystemService.
//
// It takes a *repositorys.ApiSystemRepository as a parameter and returns a *ApiSystemService.
func NewApiSystemService(repo *repositorys.ApiSystemRepository) *ApiSystemService {
    return &ApiSystemService{repository: repo}
}

func (s *ApiSystemService) GetDataService(filter types.ApiSystem) ([]*types.ApiSystem, error) {
    return s.repository.Get(filter)
}
