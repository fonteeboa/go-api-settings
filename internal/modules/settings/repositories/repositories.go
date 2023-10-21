package repositories

import (
    "github.com/jinzhu/gorm"
    "golang-api-settings/internal/modules/settings/domain"
)
type SettingsRepository struct {
    db *gorm.DB
}

// NewSettingsRepository creates a new instance of the SettingsRepository struct.
//
// It takes a *gorm.DB parameter named db and returns a *SettingsRepository.
func NewSettingsRepository(db *gorm.DB) *SettingsRepository {
    return &SettingsRepository{db: db}
}

// CreateSettings creates a new setting in the database.
//
// db: the Gorm database connection.
// setting: the setting to be created.
// error: an error if the operation fails.
func (r *SettingsRepository) CreateSettings(setting *domain.Settings) error {
    return r.db.Create(setting).Error
}


// DeleteSettingsByID deletes a Settings record from the database by its ID.
//
// Parameters:
// - db: A pointer to a gorm.DB instance representing the database connection.
// - id: The ID of the Settings record to be deleted.
//
// Returns:
// - An error if there was an issue deleting the record from the database.
func (r *SettingsRepository) DeleteSettingsByID(id uint) error {
    return r.db.Where("id = ?", id).Delete(&domain.Settings{}).Error
}


// UpdateSettings updates the settings in the database.
//
// It takes a *gorm.DB object and a *Settings object as parameters.
// It returns an error.
func (r *SettingsRepository) UpdateSettings(setting *domain.Settings) error {
    return r.db.Save(setting).Error
}


// GetData retrieves a list of settings based on the provided filter.
//
// It takes a *gorm.DB pointer as the first parameter, which represents the database connection,
// and a Settings struct as the second parameter, which contains the filter criteria.
//
// It returns a slice of Settings and an error. The slice contains the retrieved settings that match the filter,
// and the error indicates if any error occurred during the retrieval process.
func (r *SettingsRepository) GetData(filter domain.Settings) ([]*domain.Settings, error) {
    var settings []*domain.Settings
    query := r.db

    if filter.ApiKey != "" {
        query = query.Where("api_key LIKE ?", "%"+filter.ApiKey+"%")
    }
    if filter.ApiId != 0 {
        query = query.Where("api_id = ?", filter.ApiId)
    }

    if err := query.Find(&settings).Error; err != nil {
        return nil, err
    }

    return settings, nil
}


// GetDataWithJoin retrieves a list of Settings from the database based on the provided filter.
//
// It takes in a *gorm.DB object representing the database connection and a Settings object representing the filter.
// The function returns a slice of Settings and an error object.
func (r *SettingsRepository) GetDataWithJoin(filter domain.Settings) ([]*domain.Settings, error) {
    var settings []*domain.Settings
    query := r.db.Preload("ApiSystem")

    if filter.ApiKey != "" {
        query = query.Where("api_key LIKE ?", "%"+filter.ApiKey+"%")
    }
    if filter.ApiId != 0 {
        query = query.Where("api_id = ?", filter.ApiId)
    }

    if err := query.Find(&settings).Error; err != nil {
        return nil, err
    }

    return settings, nil
}
