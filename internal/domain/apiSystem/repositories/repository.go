package repositorys

import (
    "github.com/vingarcia/ksql"
    "context"
    "golang-api-settings/internal/domain/apiSystem/types"
)

type ApiSystemRepository struct {
    db ksql.DB
}

func NewApiSystemRepository(db ksql.DB) *ApiSystemRepository {
    return &ApiSystemRepository{db: db}
}


func (r *ApiSystemRepository) Get(filter types.ApiSystem) (settings []*types.ApiSystem, err error) {
    // Crie um contexto
    ctx := context.Background()

    // Construir a consulta SQL base
    query := "SELECT * FROM apiSystem WHERE 1 = 1"
    args := []interface{}{}

    // Verificar se um filtro foi fornecido e adicionar condições à consulta
    if filter.Name != "" {
        query += " AND name LIKE ?"
        args = append(args, "%"+filter.Name+"%")
    }

    if filter.Value != "" {
        query += " AND value LIKE ?"
        args = append(args, "%"+filter.Value+"%")
    }

    // Verificar se o filtro do ID foi fornecido
    if filter.ID != 0 {
        query += " AND id = ?"
        args = append(args, filter.ID)
    }
    
    // Executar a consulta preparada
    err = r.db.Query(ctx, &settings, query , args)

    if err != nil {
		return nil, err
	}

    return settings, nil
}
