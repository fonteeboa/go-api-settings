package database

import (
    "github.com/jinzhu/gorm"
    "fmt"
)

func (dbConn *GormDBHandler) Select(tableName string, params QueryParams) (*gorm.DB, error) {
    query := dbConn.DB.Table(tableName)

    // Aplicar joins, se forem fornecidos
    if len(params.Joins) > 0 {
        for _, join := range params.Joins {
            query = query.Joins(join.JoinType, join.Table).Where(join.Condition)
        }
    }

    // Aplicar filtros, se forem fornecidos
    if len(params.Filters) > 0 {
        for _, filter := range params.Filters {
            switch filter.Operator {
            case "eq":
                query = query.Where(filter.Field+" = ?", filter.Value)
            case "lt":
                query = query.Where(filter.Field+" < ?", filter.Value)
            case "gt":
                query = query.Where(filter.Field+" > ?", filter.Value)
            case "like":
                query = query.Where(filter.Field+" LIKE ?", "%"+filter.Value.(string)+"%")
            default:
                return nil, fmt.Errorf("Operador de filtro não suportado: %s", filter.Operator)
            }
        }
    }

    return query, nil
}



func (dbConn *GormDBHandler) Open() error {
    return dbConn.DB.DB().Ping()
}

func (dbConn *GormDBHandler) Close() error {
    return dbConn.DB.Close()
}

func (dbConn *GormDBHandler) AutoMigrate(value interface{}) error {
    return dbConn.DB.AutoMigrate(value).Error
}

func (dbConn *GormDBHandler) Create(value interface{}) error {
    return dbConn.DB.Create(value).Error
}

func (dbConn *GormDBHandler) Find(out interface{}, where ...interface{}) error {
    return dbConn.DB.Find(out, where...).Error
}

func (dbConn *GormDBHandler) Save(value interface{}) error {
    return dbConn.DB.Save(value).Error
}

// DeleteGeneric exclui registros com base em um modelo e critérios personalizados.
func (dbConn *GormDBHandler) Delete(model interface{}, where ...interface{}) error {
    result := dbConn.DB.Where(where).Delete(model)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (dbConn *GormDBHandler) RawSQL(sql string, values ...interface{}) error {
    return dbConn.DB.Raw(sql, values...).Error
}

func (dbConn *GormDBHandler) Exec(sql string, values ...interface{}) error {
    return dbConn.DB.Exec(sql, values...).Error
}

func (dbConn *GormDBHandler) Model(value interface{}) *GormDBHandler {
    dbConn.DB = dbConn.DB.Model(value)
    return dbConn
}