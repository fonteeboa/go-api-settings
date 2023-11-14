package domain

import (
	databaseDomain "golang-api-settings/internal/infra/database/exportDomain"
)

var Constraints = []databaseDomain.ForeignKeyConfig{
	{
		Model:          &ApiSystem{},
		ForeignKeyName: "api_id",
		References:     "settings(id)",
		OnUpdate:       "CASCADE",
		OnDelete:       "SET NULL",
	},
}
