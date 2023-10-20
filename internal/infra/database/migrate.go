package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
    settingsModels "golang-api-settings/internal/modules/settings/domain"
    apiSystemModels "golang-api-settings/internal/modules/apiSystem/domain"

)

func MigrateAllModels(db *gorm.DB) {
	models := []interface{}{ // Adicione todos os modelos aqui
        &apiSystemModels.ApiSystem{},
        &settingsModels.Settings{},
		// Adicione todos os seus modelos
	}

	for _, model := range models {
		if err := db.AutoMigrate(model).Error; err != nil {
			fmt.Printf("Erro ao migrar modelo %T: %v\n", model, err)
		} else {
			fmt.Printf("Modelo %T migrado com sucesso.\n", model)
		}
	}
}
