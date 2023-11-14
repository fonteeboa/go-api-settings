package database

import (
	"fmt"
	apiSystemModels "golang-api-settings/internal/modules/apiSystem/domain"
	settingsModels "golang-api-settings/internal/modules/settings/domain"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func MigrateAllModels(db DBHandler) {
	models := []interface{}{ // Adicione todos os modelos aqui
		&settingsModels.Settings{},
		&apiSystemModels.ApiSystem{},
		// Adicione todos os seus modelos
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Modelo %T migrado com sucesso.\n", model)
		}
	}

	MigrateCreateConstraints(db)

	// Inserir dados iniciais para a tabela Settings
	InsertBaseData(db)
}

func MigrateCreateConstraints(db DBHandler) {
	allConstraints := MergeConstraints(
		apiSystemModels.Constraints,
	)
	err := db.AddForeignKeys(allConstraints)
	if err != nil {
		fmt.Printf("Erro ao adicionar chaves estrangeiras: %v\n", err)
	}

}
