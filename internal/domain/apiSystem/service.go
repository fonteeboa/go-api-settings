package apiSystem

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type ApiSystemService struct {
    repository *ApiSystemRepository
}

func NewApiSytemService(repo *ApiSystemRepository) *ApiSystemService {
    return &ApiSystemService{repository: repo}
}

func (s *ApiSystemService) GetDataService(ctx *gin.Context) {

    var filter ApiSystem
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
