package apiSystem

import (
	"golang-api-settings/internal/helpers"
	types "golang-api-settings/internal/modules/apiSystem/domain"
	"golang-api-settings/internal/modules/apiSystem/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	service *services.ApiSystemService
}

func NewApiController(service *services.ApiSystemService) *ApiController {
	return &ApiController{service: service}
}

func (c *ApiController) GetData(ctx *gin.Context) {
	var filter types.ApiSystem

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
	var filter types.ApiSystem

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

func (c *ApiController) InsertData(ctx *gin.Context) {
	var filter *types.ApiSystem

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
	var filter *types.ApiSystem

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
