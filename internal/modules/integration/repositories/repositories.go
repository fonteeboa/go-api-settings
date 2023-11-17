package repositories

import (
	database "golang-api-settings/internal/infra/database"
	databaseDomain "golang-api-settings/internal/infra/database/exportDomain"
	types "golang-api-settings/internal/modules/integration/domain"
)

type IntegrationRepository struct {
	db database.DBHandler
}

func NewIntegrationRepository(db database.DBHandler) *IntegrationRepository {
	return &IntegrationRepository{db: db}
}

// Create is a function that creates an Integration in the database.
//
// It takes an Integration pointer as a parameter and returns an error.
func (r *IntegrationRepository) Create(integration *types.Integration) error {
	return r.db.Create(integration)
}

// DeleteByID deletes an API system by its ID.
//
// Parameters:
// - id: the ID of the API system to delete.
//
// Returns:
// - error: an error if the deletion fails.
func (r *IntegrationRepository) DeleteByID(id uint) error {
	return r.db.Delete(&types.Integration{}, id)
}

// Update updates the given Integration in the IntegrationRepository.
//
// It takes an Integration pointer as a parameter and returns an error.
// Update updates the given Integration in the IntegrationRepository.
//
// It takes an Integration pointer as a parameter and returns an error.
func (r *IntegrationRepository) Update(integration *types.Integration) error {
	return r.db.Save(integration)
}

// GetData retrieves data from the Integration repository based on the provided filter.
//
// The filter parameter is an instance of the types.Integration struct, which is used to filter the data.
// The function returns a slice of types.Integration and an error.
func (r *IntegrationRepository) GetData(filter types.Integration) ([]*types.Integration, error) {
	// Criar uma lista para armazenar os filtros
	filters := []databaseDomain.Filter{}

	// Adicionar filtros específicos
	if filter.ApiKey != "" {
		filters = append(filters, databaseDomain.Filter{Field: "api_key", Operator: "like", Value: "%" + filter.ApiKey + "%"})
	}

	if filter.ApiID != 0 {
		filters = append(filters, databaseDomain.Filter{Field: "api_id", Operator: "eq", Value: filter.ApiID})
	}

	if filter.ID != 0 {
		filters = append(filters, databaseDomain.Filter{Field: "id", Operator: "eq", Value: filter.ID})
	}

	joins := []databaseDomain.Join{}
	joins = append(joins, databaseDomain.Join{JoinType: "LEFT JOIN", Table: "authorizations", Condition: "authorizations.id = integrations.api_id"})

	// Criar parâmetros com base nos filtros
	params := databaseDomain.QueryParams{
		Filters: filters,
		Joins:   joins,
	}

	// Criar a consulta
	var integration []*types.Integration

	tableName := r.db.GetTableName(&types.Integration{})
	query, errQuery := r.db.MountQuery(tableName, params)
	if errQuery != nil {
		return nil, errQuery
	}

	err := query.Preload("Authorization").Find(&integration).Error
	if err != nil {
		return nil, err
	}

	return integration, nil
}
