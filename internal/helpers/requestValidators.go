package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateAndBindJSON validates and binds JSON data from the request body to a given filter.
//
// It checks if the request body is present and not empty. If it is, the filter is set to nil and
// the function returns true, indicating that the validation passed. If the body is not empty,
// it attempts to bind the JSON data to the filter using ctx.BindJSON(). If an error occurs,
// the error is printed and a JSON response with a bad request status code is returned. The function
// then returns false, indicating that a validation error occurred. Otherwise, if the binding is
// successful, the function returns true, indicating that the validation passed successfully.
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
