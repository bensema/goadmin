package sql

import (
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
)

func SqlBuilder() *entSql.DialectBuilder {
	return entSql.Dialect(dialect.MySQL)
}
