package repositorys

import (
    "github.com/jinzhu/gorm"
    "golang-api-settings/internal/modules/apiSystem/types"
)

type ApiSystemRepository struct {
    db *gorm.DB
}

func NewApiSystemRepository(db *gorm.DB) *ApiSystemRepository {
    return &ApiSystemRepository{db: db}
}


func (r *ApiSystemRepository) Get(filter types.ApiSystem) (settings []*types.ApiSystem, err error) {

    // Construir a consulta SQL base
    query := "SELECT * FROM apiSystem WHERE 1 = 1"
    args := []interface{}{}

    // Verificar se um filtro foi fornecido e adicionar condições à consulta
    if filter.Name != "" {
        query += " AND name LIKE ?"
        args = append(args, "%"+filter.Name+"%")
    }

    if filter.ApiName != "" {
        query += " AND apiName LIKE ?"
        args = append(args, "%"+filter.ApiName+"%")
    }

    // Verificar se o filtro do ID foi fornecido
    if filter.ID != 0 {
        query += " AND id = ?"
        args = append(args, filter.ID)
    }
    
    // Executar a consulta preparada
    if err := r.db.Raw(query, args...).Scan(&settings).Error; err != nil {
        return nil, err
    }

    return settings, nil
}
