package database

import (
	"fmt"
	databaseDomain "golang-api-settings/internal/infra/database/exportDomain"
	"strings"
)

// InitialData é uma interface para dados iniciais genéricos
type InitialData interface{}

var InitialDataMap = make(map[string][]InitialData)

// InsertBaseData insere dados iniciais para todos os modelos registrados
func InsertBaseData(db DBHandler) {
	fmt.Println("Função de inserção de dados iniciais em execução.")
	AddSettingsInitialData()
	for modelName := range InitialDataMap {
		fmt.Printf("Migrando dados iniciais para o modelo %s\n", modelName)
		InsertDataForModel(db, modelName)
	}

	fmt.Println("Função de inserção de dados iniciais concluída.")
}

// InsertDataForModel insere dados iniciais para um modelo específico
func InsertDataForModel(db DBHandler, modelName string) {
	data, ok := InitialDataMap[modelName]
	if !ok {
		fmt.Printf("Não há dados iniciais para o modelo %s.\n", modelName)
		return
	}
	modelNameLowerCaseString := strings.ToLower(modelName)

	for _, eachRow := range data {

		var count int
		db.Table(modelNameLowerCaseString).Where(eachRow).Count(&count)

		if count == 0 {
			if err := db.Create(eachRow); err != nil {
				fmt.Printf("Erro ao inserir dados iniciais para o modelo %s: %v\n", modelName, err)
			} else {
				fmt.Printf("Dados iniciais para o modelo %s inseridos com sucesso.\n", modelName)
			}
		} else {
			fmt.Printf("Dados iniciais para o modelo %s já existem. Nenhuma inserção necessária.\n", modelName)
		}

	}
}

// AddInitialData adiciona dados iniciais para um modelo específico
func AddInitialData(modelName string, data InitialData) {
	InitialDataMap[modelName] = append(InitialDataMap[modelName], data)
}
func MergeConstraints(arrays ...[]databaseDomain.ForeignKeyConfig) []databaseDomain.ForeignKeyConfig {
	var result []databaseDomain.ForeignKeyConfig
	for _, array := range arrays {
		result = append(result, array...)
	}
	return result
}
