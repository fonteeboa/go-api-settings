package integration

import (
	"golang-api-settings/internal/helpers"
	types "golang-api-settings/internal/modules/integration/domain"
	"golang-api-settings/internal/modules/integration/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	service *services.IntegrationService
}

func NewApiController(service *services.IntegrationService) *ApiController {
	return &ApiController{service: service}
}

func (c *ApiController) GetData(ctx *gin.Context) {
	var filter types.Integration

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

func (c *ApiController) Delete(ctx *gin.Context) {
	var filter types.Integration

	if !helpers.ValidateAndBindJSON(ctx, &filter) {
		return // Se a validação falhar, não continue com o processamento
	}

	// Chame o serviço para obter os dados
	err := c.service.Delete(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorne os dados da resposta
	ctx.JSON(http.StatusOK, 1)
}

func (c *ApiController) DeleteAll(ctx *gin.Context) {
	var filters []types.Integration

	// Bind do JSON para um slice (array) de types.Integration
	if err := ctx.BindJSON(&filters); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := c.service.DeleteAll(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorne os dados da resposta
	ctx.JSON(http.StatusOK, 1)

	// Retorna um status OK (200) após a exclusão bem-sucedida
	ctx.JSON(http.StatusOK, gin.H{"message": "Items deleted successfully"})
}

func (c *ApiController) InsertData(ctx *gin.Context) {
	var filter *types.Integration

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

func (c *ApiController) UpdateData(ctx *gin.Context) {
	var filter *types.Integration

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
