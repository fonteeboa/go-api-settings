package database

import (
	"fmt"
	"golang-api-settings/internal/infra/database/initialData"
	apiSystemModels "golang-api-settings/internal/modules/apiSystem/domain"
	settingsModels "golang-api-settings/internal/modules/settings/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	// Inserir dados iniciais para a tabela Settings
	initialData.InsertBaseData(db)
}
