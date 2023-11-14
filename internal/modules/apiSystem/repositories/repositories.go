package repositories

import (
	database "golang-api-settings/internal/infra/database"
	databaseDomain "golang-api-settings/internal/infra/database/exportDomain"
	types "golang-api-settings/internal/modules/apiSystem/domain"
)

type ApiSystemRepository struct {
	db database.DBHandler
}

func NewApiSystemRepository(db database.DBHandler) *ApiSystemRepository {
	return &ApiSystemRepository{db: db}
}

// Create is a function that creates an ApiSystem in the database.
//
// It takes an ApiSystem pointer as a parameter and returns an error.
func (r *ApiSystemRepository) Create(apiSystem *types.ApiSystem) error {
	return r.db.Create(apiSystem)
}

// DeleteByID deletes an API system by its ID.
//
// Parameters:
// - id: the ID of the API system to delete.
//
// Returns:
// - error: an error if the deletion fails.
func (r *ApiSystemRepository) DeleteByID(id uint) error {
	return r.db.Delete(&types.ApiSystem{}, id)
}

// Update updates the given ApiSystem in the ApiSystemRepository.
//
// It takes an ApiSystem pointer as a parameter and returns an error.
// Update updates the given ApiSystem in the ApiSystemRepository.
//
// It takes an ApiSystem pointer as a parameter and returns an error.
func (r *ApiSystemRepository) Update(apiSystem *types.ApiSystem) error {
	return r.db.Save(apiSystem)
}

// GetData retrieves data from the ApiSystem repository based on the provided filter.
//
// The filter parameter is an instance of the types.ApiSystem struct, which is used to filter the data.
// The function returns a slice of types.ApiSystem and an error.
func (r *ApiSystemRepository) GetData(filter types.ApiSystem) ([]*types.ApiSystem, error) {
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
	joins = append(joins, databaseDomain.Join{JoinType: "LEFT JOIN", Table: "settings", Condition: "settings.id = api_systems.api_id"})

	// Criar parâmetros com base nos filtros
	params := databaseDomain.QueryParams{
		Filters: filters,
		Joins:   joins,
	}

	// Criar a consulta
	var apiSystem []*types.ApiSystem

	tableName := r.db.GetTableName(&types.ApiSystem{})
	query, errQuery := r.db.MountQuery(tableName, params)
	if errQuery != nil {
		return nil, errQuery
	}

	err := query.Find(&apiSystem).Error
	if err != nil {
		return nil, err
	}

	return apiSystem, nil
}
