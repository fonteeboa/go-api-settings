package apiSystem

import (
    "github.com/gin-gonic/gin"
    "net/http"
    types "golang-api-settings/internal/modules/apiSystem/domain"
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


func (c *ApiController) Delete(ctx *gin.Context) {
    var filter types.ApiSystem
    if err := ctx.BindJSON(&filter); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
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
    if err := ctx.BindJSON(&filter); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
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
