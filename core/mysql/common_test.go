package mysql

import (
	"github.com/go-xorm/xorm"
	"testing"
)

func runHandler(_orm *xorm.Session) (interface{}, SessionDoctor, error) {
	t1 := TestTrans{Id: 1}
	if _, err := _orm.InsertOne(t1); err != nil {
		return nil, 1, err
	}
	return nil, 0, nil
}

func TestCommitRollBack(t *testing.T) {
	c := Config{
		ShowSql: true,
		Master:  "root:123456@tcp(127.0.0.1)/virtual_coin?charset=utf8mb4&parseTime=true&loc=Local",
	}
	conn, _ := NewConn(&c)
	dao := NewBaseDao(conn)
	session := dao.NewSession().Session()
	defer session.Close()
	WarpSession(session, runHandler)
}

func TestNewSessionWrapper(t *testing.T) {
	c := Config{
		ShowSql: true,
		Master:  "root:123456@tcp(127.0.0.1)/virtual_coin?charset=utf8mb4&parseTime=true&loc=Local",
	}
	conn, _ := NewConn(&c)
	dao := NewBaseDao(conn)
	s := NewSessionWrapper(dao)
	err := s.Execute(func(session *xorm.Session) error {
		t1 := TestTrans{Id: 1}
		if _, err := session.InsertOne(t1); err != nil {
			return err
		}
		return nil
	}).Execute(func(session *xorm.Session) error {
		t2 := TestTrans{Id: 2}
		if _, err := session.InsertOne(t2); err != nil {
			return err
		}
		return nil
	}).Execute(func(session *xorm.Session) error {
		t3 := TestTrans{Id: 3}
		if _, err := session.InsertOne(t3); err != nil {
			return err
		}
		return nil
	}).Commit()
	if err != nil {
		t.Fatal(err)
	}

}
