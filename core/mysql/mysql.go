package mysql

import (
	"AdServerCollector/conf"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hashicorp/go-hclog"
	_ "github.com/lib/pq"
	"log"
	"os"
	"xorm.io/core"
)

// DBInterface 数据库接口
type DBInterface interface {
	xorm.EngineInterface
}

func NewConn(c *Config, logFile string) (DBInterface, error) {
	var logger hclog.Logger
	if logFile != "" {
		logFn, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("open file error !")
			return nil, err
		}
		logger = hclog.New(&hclog.LoggerOptions{
			Name:   conf.Config.Global.AppName,
			Level:  hclog.LevelFromString("DEBUG"),
			Output: logFn,
		})
	} else {
		logger = hclog.Default()
	}

	if c.UseMasterSlave {
		return NewMSConn(c, logger)
	}
	return NewSingleConn(c, logger)
}

// NewSingleConn 初始化数据库连接
// mysql fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",user,pwd,host,db,charset)
// postgres fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s",
//
//	host, port, user, name, pass, sslMode, SslCert, SslKey, SslRootCert)
func NewSingleConn(c *Config, logger hclog.Logger) (DBInterface, error) {
	if nil == c || "" == c.Master {
		return nil, errors.New("config or config.Url can not be null")
	}
	var conn *xorm.Engine
	var err error
	switch c.DbType {
	case "mysql":
		conn, err = xorm.NewEngine("mysql", c.Master)
	case "postgres":
		conn, err = xorm.NewEngine("postgres", c.Master)
	default:
		conn, err = xorm.NewEngine("mysql", c.Master)
	}
	if nil != err || nil == conn {
		log.Println("failed to initializing db connection:", err)
		return nil, err
	}
	conn.SetMapper(core.GonicMapper{})
	conn.SetLogger(NewSQLLogger(logger, c.ShowSql))
	conn.ShowSQL(c.ShowSql)
	conn.SetMaxIdleConns(c.MaxIdle)
	conn.SetMaxOpenConns(c.MaxConn)
	return conn, nil
}

// NewMSConn 初始化主从数据库连接, master不能为空，slaves可以为空
func NewMSConn(c *Config, logger hclog.Logger) (DBInterface, error) {
	if nil == c || "" == c.Master {
		return nil, errors.New("config or config.Url can not be null")
	}
	conns := make([]string, len(c.Slaves)+1)
	conns[0] = c.Master
	for i, v := range c.Slaves {
		conns[i+1] = v
		if "" == v {
			return nil, errors.New("config or config.Url can not be null")
		}
	}

	var group *xorm.EngineGroup
	var err error
	switch c.DbType {
	case "mysql":
		group, err = xorm.NewEngineGroup("mysql", c.Master)
	case "postgres":
		group, err = xorm.NewEngineGroup("postgres", c.Master)
	default:
		group, err = xorm.NewEngineGroup("mysql", c.Master)
	}

	if nil != err || nil == group {
		log.Printf("failed to initializing db connection: %s\n", err)
		return nil, err
	}

	group.SetMapper(core.GonicMapper{})
	group.SetLogger(NewSQLLogger(logger, c.ShowSql))
	group.ShowSQL(c.ShowSql)
	group.SetMaxIdleConns(c.MaxIdle)
	group.SetMaxOpenConns(c.MaxConn)
	return group, nil
}

var dbConn DBInterface

func init() {
	_conf, err := json.Marshal(conf.Config.Mysql)
	if err != nil {
		panic(err)
	}
	var dbConfig Config
	err = json.Unmarshal(_conf, &dbConfig)
	dbConfig.ShowSql = true
	if err != nil {
		panic(err)
	}
	dbConn, _ = NewConn(&dbConfig, conf.Config.Mysql.LogFile)
}

func GetMySqlDB() DBInterface {
	return dbConn
}

func GetDao() BaseDao {
	return NewBaseDao(dbConn)
}
