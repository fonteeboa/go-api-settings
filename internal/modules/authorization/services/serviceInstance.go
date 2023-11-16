package services

import (
	"golang-api-settings/internal/modules/authorization/repositories"
)

type AuthorizationService struct {
	repositories *repositories.AuthorizationRepository
}

func NewAuthorizationService(repo *repositories.AuthorizationRepository) *AuthorizationService {
	return &AuthorizationService{repositories: repo}
}
