package apiSystem

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "golang-api-settings/internal/modules/apiSystem/types"
    "golang-api-settings/internal/modules/apiSystem/services"
)

type ApiController struct {
    service *services.ApiSystemService
}

func NewApiController(service *services.ApiSystemService) *ApiController {
    return &ApiController{service: service}
}

func (c *ApiController) GetData(ctx *gin.Context) {
    var filter types.ApiSystem
    if err := ctx.BindJSON(&filter); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Chame o serviço para obter os dados
    response, err := c.service.GetDataService(filter)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Retorne os dados da resposta
    ctx.JSON(http.StatusOK, response)
}
