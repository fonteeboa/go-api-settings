package settings

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type SettingsService struct {
    repository *SettingsRepository
}

func NewSettingsService(repo *SettingsRepository) *SettingsService {
    return &SettingsService{repository: repo}
}

func (s *SettingsService) GetDataService(ctx *gin.Context) {

    var filter Settings
    if err := ctx.BindJSON(&filter); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Chame o serviço para obter os dados
    response, err := s.repository.Get(filter)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return 
    }

    // Retorne os dados da resposta
    ctx.JSON(http.StatusOK, response)

}
