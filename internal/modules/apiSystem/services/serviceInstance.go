package services

import (
    "golang-api-settings/internal/modules/apiSystem/repositories"
)

type ApiSystemService struct {
    repositories *repositories.ApiSystemRepository
}

// NewApiSystemService creates a new instance of the ApiSystemService struct.
// It takes a pointer to an ApiSystemRepository as a parameter and returns a pointer to the ApiSystemService.
func NewApiSystemService(repo *repositories.ApiSystemRepository) *ApiSystemService {
    return &ApiSystemService{repositories: repo}
}
