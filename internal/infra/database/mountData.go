package database

import (
	authorizationModels "golang-api-settings/internal/modules/authorization/domain"
)

func AddSettingsInitialData() {
	for _, data := range authorizationModels.AuthorizationInitialData {
		AddInitialData("authorizations", data)
	}
}
