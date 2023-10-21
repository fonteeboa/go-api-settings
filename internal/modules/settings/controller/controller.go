package apiSystem

import (
    "github.com/gin-gonic/gin"
    "net/http"
    types "golang-api-settings/internal/modules/settings/domain"
    "golang-api-settings/internal/modules/settings/services"
)

type settingsController struct {
    service *services.SettingsService
}

func NewSettingsController(service *services.SettingsService) *settingsController {
    return &settingsController{service: service}
}

func (c *settingsController) GetData(ctx *gin.Context) {
    var filter types.Settings
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
