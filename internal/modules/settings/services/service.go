package services

import (
    "golang-api-settings/internal/modules/settings/types"
    "golang-api-settings/internal/modules/settings/repositories"
)


type SettingsService struct {
    repositories *repositories.SettingsRepository
}

func NewSettingsService(repo *repositories.SettingsRepository) *SettingsService {
    return &SettingsService{repositories: repo}
}

func (s *SettingsService) GetDataService(filter types.Settings) ([]*types.Settings, error) {
    return s.repositories.Get(filter)
}
