package apiSystem

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "golang-api-settings/internal/domain/apiSystem/types"
    "golang-api-settings/internal/domain/apiSystem/repositorys"
)

type ApiSystemService struct {
    repository *repositorys.ApiSystemRepository
}

func NewApiSytemService(repo *repositorys.ApiSystemRepository) *ApiSystemService {
    return &ApiSystemService{repository: repo}
}

func (s *ApiSystemService) GetDataService(ctx *gin.Context) {

    var filter types.ApiSystem
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
