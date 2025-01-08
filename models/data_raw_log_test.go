package models

import (
	"AdServerCollector/core/mysql"
	"testing"
)

func TestSliceDateTable(t *testing.T) {
	c := mysql.Config{
		ShowSql: true,
		Master:  "root:123456@tcp(127.0.0.1)/ad-aws?charset=utf8mb4&parseTime=true&loc=Local",
	}
	conn, _ := mysql.NewConn(&c)
	m := DataRawLog{}
	b, err := conn.Table(m.GetSliceDateDayTable()).Exist()
	if err != nil {
		t.Error(err)
	}
	if !b {
		err := conn.Table(m.GetSliceDateDayTable()).CreateTable(m)
		if err != nil {
			t.Error(err.Error())
		}
	}
}
