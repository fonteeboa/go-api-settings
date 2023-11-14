package database

import (
	settingsModels "golang-api-settings/internal/modules/settings/domain"
)

func AddSettingsInitialData() {
	for _, data := range settingsModels.SettingsInitialData {
		AddInitialData("Settings", data)
	}
}
