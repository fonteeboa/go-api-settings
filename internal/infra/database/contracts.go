package database

import (
	"context"

	"github.com/vingarcia/ksql"
)

type Database interface {
	Insert(ctx context.Context, table ksql.Table, record interface{}) error
	QueryOne(ctx context.Context, record interface{}, query string, params ...interface{}) error
	Patch(ctx context.Context, table ksql.Table, record interface{}) error
	Query(ctx context.Context, records interface{}, query string, params ...interface{}) error
	Delete(ctx context.Context, table ksql.Table, idOrRecord interface{}) error
}
