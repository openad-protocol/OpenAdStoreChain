package mysql

import (
	"testing"
)

func TestEntity_Create(t *testing.T) {
	c := Config{
		ShowSql: true,
		Master:  "root:123456@tcp(127.0.0.1)/virtual_coin?charset=utf8mb4&parseTime=true&loc=Local",
	}
	conn, _ := NewConn(&c)
	createTable := `create table if not exists test_trans(id int)engine=innodb`
	conn.Query(createTable)
}
