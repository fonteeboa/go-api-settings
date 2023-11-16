package repositories

import (
	"golang-api-settings/internal/infra/database"
	databaseDomain "golang-api-settings/internal/infra/database/exportDomain"
	"golang-api-settings/internal/modules/authorization/domain"

	"gorm.io/gorm"
)

type AuthorizationRepository struct {
	db database.DBHandler
}

// NewAuthorizationRepository creates a new instance of the AuthorizationRepository struct.
//
// It takes a *gorm.DB parameter named db and returns a *AuthorizationRepository.
func NewAuthorizationRepository(db database.DBHandler) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

// CreateAuthorization creates a new setting in the database.
//
// db: the Gorm database connection.
// setting: the setting to be created.
// error: an error if the operation fails.
func (r *AuthorizationRepository) CreateAuthorization(setting *domain.Authorization) error {
	return r.db.Create(setting)
}

// DeleteAuthorizationByID deletes a Authorization record from the database by its ID.
//
// Parameters:
// - db: A pointer to a gorm.DB instance representing the database connection.
// - id: The ID of the Authorization record to be deleted.
//
// Returns:
// - An error if there was an issue deleting the record from the database.
func (r *AuthorizationRepository) DeleteAuthorizationByID(id uint) error {
	// Crie uma instância do modelo com o campo ID especificado
	settingsToDelete := domain.Authorization{Model: gorm.Model{ID: id}}
	// Use a função Delete para excluir o registro
	if err := r.db.Delete(&settingsToDelete); err != nil {
		// Se ocorrer um erro durante a exclusão, retorne o erro
		return err
	}

	// Operação de exclusão bem-sucedida
	return nil
}

// UpdateAuthorization updates the settings in the database.
//
// It takes a *gorm.DB object and a *Authorization object as parameters.
// It returns an error.
func (r *AuthorizationRepository) UpdateAuthorization(setting *domain.Authorization) error {
	return r.db.Save(setting)
}

// GetData retrieves a list of settings based on the provided filter.
//
// It takes a *gorm.DB pointer as the first parameter, which represents the database connection,
// and a Authorization struct as the second parameter, which contains the filter criteria.
//
// It returns a slice of Authorization and an error. The slice contains the retrieved settings that match the filter,
// and the error indicates if any error occurred during the retrieval process.
func (r *AuthorizationRepository) GetData(filter domain.Authorization) ([]*domain.Authorization, error) {
	var settings []*domain.Authorization
	filters := []databaseDomain.Filter{}

	// Adicionar filtros específicos
	if filter.Name != "" {
		filters = append(filters, databaseDomain.Filter{Field: "name", Operator: "like", Value: "%" + filter.Name + "%"})
	}

	if filter.ID != 0 {
		filters = append(filters, databaseDomain.Filter{Field: "id", Operator: "eq", Value: filter.ID})
	}

	// Criar parâmetros com base nos filtros
	params := databaseDomain.QueryParams{
		Filters: filters,
		// Pode adicionar Joins aqui se necessário
	}

	tableName := r.db.GetTableName(&domain.Authorization{})
	query, errQuery := r.db.MountQuery(tableName, params)
	if errQuery != nil {
		return nil, errQuery
	}

	err := query.Find(&settings).Error
	if err != nil {
		return nil, err
	}

	return settings, nil
}
