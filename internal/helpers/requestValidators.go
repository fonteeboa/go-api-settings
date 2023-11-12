package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Função para validar o corpo da requisição e fazer BindJSON
func ValidateAndBindJSON(ctx *gin.Context, filter interface{}) bool {
	// Verifica se o corpo está presente e não está vazio
	if ctx.Request.Body == nil || ctx.Request.ContentLength == 0 {
		// Corpo ausente ou vazio, define filter como um valor vazio
		filter = nil
		return true // Indica que a validação passou, pois não há corpo para validar
	}

	// Faz o BindJSON apenas se o corpo estiver presente
	if err := ctx.BindJSON(filter); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "common.body.error"})
		return false // Indica que houve um erro de validação
	}

	return true // Indica que a validação passou com sucesso
}
