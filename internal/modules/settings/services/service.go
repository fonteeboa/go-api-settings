package services

import (
    "golang-api-settings/internal/modules/settings/domain"
    "golang-api-settings/internal/modules/settings/repositories"
)


type SettingsService struct {
    repositories *repositories.SettingsRepository
}

func NewSettingsService(repo *repositories.SettingsRepository) *SettingsService {
    return &SettingsService{repositories: repo}
}

func (s *SettingsService) GetDataService(filter domain.Settings) ([]*domain.Settings, error) {
    return s.repositories.GetData(filter)
}
