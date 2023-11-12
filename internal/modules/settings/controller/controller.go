package apiSystem

import (
	"golang-api-settings/internal/helpers"
	types "golang-api-settings/internal/modules/settings/domain"
	"golang-api-settings/internal/modules/settings/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type settingsController struct {
	service *services.SettingsService
}

func NewSettingsController(service *services.SettingsService) *settingsController {
	return &settingsController{service: service}
}

func (c *settingsController) GetData(ctx *gin.Context) {
	var filter types.Settings

	if !helpers.ValidateAndBindJSON(ctx, &filter) {
		return // Se a validação falhar, não continue com o processamento
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

func (c *settingsController) InsertData(ctx *gin.Context) {
	var filter *types.Settings

	if !helpers.ValidateAndBindJSON(ctx, &filter) {
		return // Se a validação falhar, não continue com o processamento
	}

	// Chame o serviço para obter os dados
	response, err := c.service.InsertData(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorne os dados da resposta
	ctx.JSON(http.StatusOK, response)
}

func (c *settingsController) UpdateData(ctx *gin.Context) {
	var filter *types.Settings

	if !helpers.ValidateAndBindJSON(ctx, &filter) {
		return // Se a validação falhar, não continue com o processamento
	}

	// Chame o serviço para obter os dados
	response, err := c.service.UpdateData(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorne os dados da resposta
	ctx.JSON(http.StatusOK, response)
}
