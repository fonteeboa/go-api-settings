package repositories

import (
    "gorm.io/gorm"
    "golang-api-settings/internal/modules/apiSystem/types"
)

type ApiSystemRepository struct {
    db *gorm.DB
}

func NewApiSystemRepository(db *gorm.DB) *ApiSystemRepository {
    return &ApiSystemRepository{db: db}
}


// Create is a function that creates an ApiSystem in the database.
//
// It takes an ApiSystem pointer as a parameter and returns an error.
func (r *ApiSystemRepository) Create(apiSystem *types.ApiSystem) error {
    return r.db.Create(apiSystem).Error
}


// DeleteByID deletes an API system by its ID.
//
// Parameters:
// - id: the ID of the API system to delete.
//
// Returns:
// - error: an error if the deletion fails.
func (r *ApiSystemRepository) DeleteByID(id uint) error {
    return r.db.Delete(&types.ApiSystem{}, id).Error
}


// Update updates the given ApiSystem in the ApiSystemRepository.
//
// It takes an ApiSystem pointer as a parameter and returns an error.
// Update updates the given ApiSystem in the ApiSystemRepository.
//
// It takes an ApiSystem pointer as a parameter and returns an error.
func (r *ApiSystemRepository) Update(apiSystem *types.ApiSystem) error {
    return r.db.Save(apiSystem).Error
}


// GetData retrieves data from the ApiSystem repository based on the provided filter.
//
// The filter parameter is an instance of the types.ApiSystem struct, which is used to filter the data.
// The function returns a slice of types.ApiSystem and an error.
func (r *ApiSystemRepository) GetData(filter types.ApiSystem) ([]types.ApiSystem, error) {
    var apiSystem []types.ApiSystem
    query := r.db.Model(&types.ApiSystem{})

    if filter.Name != "" {
        query = query.Where("name LIKE ?", "%"+filter.Name+"%")
    }
    if filter.ID != 0 { // Compare com 0 em vez de ""
        query = query.Where("ID = ?", filter.ID)
    }

    if err := query.Find(&apiSystem).Error; err != nil {
        return nil, err
    }
    return apiSystem, nil
}
