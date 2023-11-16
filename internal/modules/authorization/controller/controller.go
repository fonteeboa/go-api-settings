package apiSystem

import (
	"golang-api-settings/internal/helpers"
	types "golang-api-settings/internal/modules/authorization/domain"
	"golang-api-settings/internal/modules/authorization/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authorizationController struct {
	service *services.AuthorizationService
}

func NewAuthorizationController(service *services.AuthorizationService) *authorizationController {
	return &authorizationController{service: service}
}

func (c *authorizationController) GetData(ctx *gin.Context) {
	var filter types.Authorization

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

func (c *authorizationController) InsertData(ctx *gin.Context) {
	var filter *types.Authorization

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

func (c *authorizationController) UpdateData(ctx *gin.Context) {
	var filter *types.Authorization

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
