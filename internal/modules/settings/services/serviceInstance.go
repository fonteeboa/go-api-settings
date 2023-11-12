package services

import (
	"golang-api-settings/internal/modules/settings/repositories"
)

type SettingsService struct {
	repositories *repositories.SettingsRepository
}

func NewSettingsService(repo *repositories.SettingsRepository) *SettingsService {
	return &SettingsService{repositories: repo}
}
