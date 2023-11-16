package database

import (
	"fmt"
	authorizationModels "golang-api-settings/internal/modules/authorization/domain"
	integrationModels "golang-api-settings/internal/modules/integration/domain"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func MigrateAllModels(db DBHandler) {
	models := []interface{}{ // Adicione todos os modelos aqui
		&authorizationModels.Authorization{},
		&integrationModels.Integration{},
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
	InsertBaseData(db)
}

func MigrateCreateConstraints(db DBHandler) {
	allConstraints := MergeConstraints(
		integrationModels.Constraints,
	)
	err := db.AddForeignKeys(allConstraints)
	if err != nil {
		fmt.Printf("Erro ao adicionar chaves estrangeiras: %v\n", err)
	}

}
