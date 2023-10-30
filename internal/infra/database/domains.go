package database

import (
    "github.com/jinzhu/gorm"
)

// DBHandler é uma interface para manipular o banco de dados Postgres com GORM.
type DBHandler interface {
    Open() error
    Close() error
    AutoMigrate(value interface{}) error
    Create(value interface{}) error
    Find(out interface{}, where ...interface{}) error
    Save(value interface{}) error
    Delete(model interface{}, where ...interface{}) error
    RawSQL(sql string, values ...interface{}) error
    Exec(sql string, values ...interface{}) error
    Model(value interface{}) DBHandler
	Select(tableName string, params QueryParams) (*gorm.DB, error)
}

// GormDBHandler é uma implementação da interface DBHandler usando GORM.
type GormDBHandler struct {
    DB *gorm.DB
}

// QueryParams é uma estrutura para representar os parâmetros de consulta dinâmica.
type QueryParams struct {
    Filters []Filter
    Joins   []Join
}

// Filter representa um filtro para a consulta.
type Filter struct {
    Field    string
    Operator string
    Value    interface{}
}

// Join representa uma cláusula JOIN para a consulta.
type Join struct {
    Table      string
    Condition  string
    JoinType   string // "INNER JOIN", "LEFT JOIN", etc.
}
