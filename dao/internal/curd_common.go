package internal

import (
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

const (
	SqlPageList  = "page_list"
	SqlPageCount = "page_count"
	SqlFind      = "find"
)

func sqlBuilder() *entsql.DialectBuilder {
	return entsql.Dialect(dialect.MySQL)
}

func checkInStr(dis []string, k string) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}
