package database

import (
	"golang-api-settings/internal/infra/database/exportDomain"

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
	Model(value interface{}) *GormDBHandler
	MountQuery(tableName string, params exportDomain.QueryParams) (*gorm.DB, error)
	Table(name string) *gorm.DB
	GetTableName(value interface{}) string
	AddForeignKeys(foreignKeys []exportDomain.ForeignKeyConfig) error
}

// GormDBHandler é uma implementação da interface DBHandler usando GORM.
// GormDBHandler implementa a interface DBHandler
type GormDBHandler struct {
	DB *gorm.DB
}
