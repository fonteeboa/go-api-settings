package services

import (
    "golang-api-settings/internal/modules/apiSystem/types"
    "golang-api-settings/internal/modules/apiSystem/repositories"
)

type ApiSystemService struct {
    repositories *repositories.ApiSystemRepository
}

// NewApiSystemService creates a new instance of the ApiSystemService.
//
// It takes a *repositories.ApiSystemRepository as a parameter and returns a *ApiSystemService.
func NewApiSystemService(repo *repositories.ApiSystemRepository) *ApiSystemService {
    return &ApiSystemService{repositories: repo}
}

func (s *ApiSystemService) GetDataService(filter types.ApiSystem) ([]*types.ApiSystem, error) {
    return s.repositories.Get(filter)
}
