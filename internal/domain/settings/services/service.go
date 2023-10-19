package services

import (
    "golang-api-settings/internal/domain/settings/types"
    "golang-api-settings/internal/domain/settings/repositories"
)


type SettingsService struct {
    repository *repositorys.SettingsRepository
}

func NewSettingsService(repo *repositorys.SettingsRepository) *SettingsService {
    return &SettingsService{repository: repo}
}

func (s *SettingsService) GetDataService(filter types.Settings) ([]*types.Settings, error) {
    return s.repository.Get(filter)
}
