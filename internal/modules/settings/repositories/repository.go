package repositorys

import (
    "github.com/jinzhu/gorm"
    "golang-api-settings/internal/modules/settings/types"
)

type SettingsRepository struct {
    db *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) *SettingsRepository {
    return &SettingsRepository{db: db}
}

func (r *SettingsRepository) Get(filter types.Settings) ([]*types.Settings, error) {

    // Construir a consulta SQL base
    query := "SELECT * FROM apiSettings WHERE 1 = 1"
    args := []interface{}{}

    // Verificar se um filtro foi fornecido e adicionar condições à consulta
    if filter.Name != "" {
        query += " AND name LIKE ?"
        args = append(args, "%"+filter.Name+"%")
    }

    if filter.Key != "" {
        query += " AND Key LIKE ?"
        args = append(args, "%"+filter.Key+"%")
    }

    // Verificar se o filtro do ID foi fornecido
    if filter.ID != 0 {
        query += " AND id = ?"
        args = append(args, filter.ID)
    }

    var settings []*types.Settings
    
    // Executar a consulta preparada
    if err := r.db.Raw(query, args...).Scan(&settings).Error; err != nil {
        return nil, err
    }

    return settings, nil
}
