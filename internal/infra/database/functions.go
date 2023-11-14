package database

import (
	"fmt"

	"golang-api-settings/internal/infra/database/exportDomain"

	"github.com/jinzhu/gorm"
)

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

func (dbConn *GormDBHandler) GetTableName(value interface{}) string {
	return dbConn.DB.NewScope(value).TableName()
}

func (dbConn *GormDBHandler) Table(name string) *gorm.DB {
	return dbConn.DB.Table(name)
}

func (dbConn *GormDBHandler) AddForeignKey(model interface{}, foreignKey, references, onUpdate, onDelete string) error {
	if err := dbConn.DB.Model(model).AddForeignKey(foreignKey, references, onUpdate, onDelete).Error; err != nil {
		return err
	}
	return nil
}

func (dbConn *GormDBHandler) AddForeignKeys(configs []exportDomain.ForeignKeyConfig) error {
	for _, config := range configs {
		if err := dbConn.DB.Model(config.Model).
			AddForeignKey(config.ForeignKeyName, config.References, config.OnUpdate, config.OnDelete).
			Error; err != nil {
			return err
		}
	}
	return nil
}

/*
MountQuery mounts a query on the specified table in the GormDBHandler.

It takes the tableName string and params QueryParams as parameters.
It returns a *gorm.DB and an error.

	Example: MountQuery("users", QueryParams{
	    Filters: []Filter{
	        {Field: "users.name", Operator: "eq", Value: "John"},
	        {Field: "roles.role_name", Operator: "eq", Value: "Admin"},
	    },
	    Joins: []Join{
	        {JoinType: "INNER JOIN", Table: "roles", Condition: "users.role_id = roles.id"},
	    },
	}
*/
func (dbConn *GormDBHandler) MountQuery(tableName string, params exportDomain.QueryParams) (*gorm.DB, error) {
	query := dbConn.DB.Table(tableName)

	// Aplicar joins, se forem fornecidos
	if len(params.Joins) > 0 {
		for _, join := range params.Joins {
			query = query.Joins(fmt.Sprintf("%s %s ON %s", join.JoinType, join.Table, join.Condition))
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
				return nil, fmt.Errorf("operador de filtro não suportado: %s", filter.Operator)
			}
		}
	}

	// Depois de construir a query, você pode obter a expressão SQL
	// dbConn.DB.LogMode(true)
	//sql := query.QueryExpr()
	//fmt.Println("Expressão SQL:", sql)

	return query, nil
}
