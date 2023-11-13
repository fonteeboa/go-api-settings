package initialData

import (
	settingsModels "golang-api-settings/internal/modules/settings/domain"
	"reflect"
)

func AddSettingsInitialData() {
	for _, data := range settingsModels.SettingsInitialData {
		AddInitialData("Settings", data)
	}
}

func getModelType(modelName string) reflect.Type {
	switch modelName {
	case "Settings":
		return reflect.TypeOf(&settingsModels.Settings{})
	// Adicione mais casos conforme necessário
	default:
		return nil
	}
}
