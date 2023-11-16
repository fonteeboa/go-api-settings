package domain

import (
	databaseDomain "golang-api-settings/internal/infra/database/exportDomain"
)

var Constraints = []databaseDomain.ForeignKeyConfig{
	{
		Model:          &Integration{},
		ForeignKeyName: "api_id",
		References:     "authorizations(id)",
		OnUpdate:       "CASCADE",
		OnDelete:       "SET NULL",
	},
}
