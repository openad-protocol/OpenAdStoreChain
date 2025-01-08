package conf

import (
	"AdServerCollector/utils"
	"crypto/tls"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/snowflake"
	"github.com/fatih/color"
)

type TonNetworkConfig struct {
	RPCUrl  string `toml:"rpc_url"`
	ChainID string `toml:"chain_id"`
	Seed    string `toml:"seed"`
}

type appCfg struct {
	Redis struct {
		Addrs        []string      `json:"addrs" toml:"addrs"`       //集群地址
		Password     string        `json:"password"`                 //密码
		KeyPrefix    string        `json:"key_prefix"`               //key前缀
		DbIndex      int           `json:"db_index" toml:"db_index"` //数据索引
		DialTimeout  time.Duration `json:"dial_timeout"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
		ReadOnly     bool          `json:"read_only"`
		// PoolSize applies per cluster node and not for the whole cluster.
		PoolSize           int           `json:"pool_size"`
		PoolTimeout        time.Duration `json:"pool_timeout"`
		IdleTimeout        time.Duration `json:"idle_timeout"`
		IdleCheckFrequency time.Duration `json:"idle_check_frequency"`
		UseCluster         bool          `json:"use_cluster"`
		TLSConfig          *tls.Config   `json:"tls_config"`
	} `json:"redis" tom:"redis"`

	Queue struct {
		Url   string `json:"url" tom:"url"`
		Topic string `json:"topic" tom:"topic"`
	} `json:"queue"`
	Mysql struct {
		ShowSql        bool     `json:"show_sql"`
		MaxIdle        int      `json:"max_idle"`
		MaxConn        int      `json:"max_conn"`
		Master         string   `json:"master"`
		Slaves         []string `json:"slaves"`
		UseMasterSlave bool     `json:"use_master_slave"`
		DbType         string   `json:"db_type"`
		LogFile        string   `json:"log_file" toml:"log_file"`
		DbVersion      string   `json:"db_version" toml:"db_version"`
	} `json:"mysql"`
	Nats struct {
		Host         string   `json:"host" tom:"host"`
		Port         int      `json:"port" tom:"port"`
		Subjects     []string `json:"subjects" tom:"subjects"`
		ConsumerName string   `json:"consumer_name" toml:"consumer_name"`
	}
	Mongo struct {
		Ip   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"mongo"`
	Global struct {
		Version string `json:"version" toml:"version"`
		AppName string `json:"app_name" toml:"app_name"`
	} `json:"global"`
	TonNetwork struct {
		Testnet TonNetworkConfig `toml:"testnet"`
		Mainnet TonNetworkConfig `toml:"mainnet"`
	} `toml:"ton_network"`
	Hash struct {
		GetAd      bool `json:"get_ad"`
		LogInfo    bool `json:"log_info"`
		ClickInfo  bool `json:"click_info"`
		TracerHash bool `json:"tracer_hash"`
		AdInCall   bool `json:"ad_in_call"`
		GetAdMiss  bool `json:"get_ad_miss"`
	} `toml:"hash"`
}

// user:password@tcp(host:port)/dbname?charset=utf8mb4
// 用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8mb4
const _dsn = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local"

const _addr = "%s:%d"

const _mongo_addr = "mongodb://%s:%d"

var (
	defConfig  = "./conf/dev.toml"  //配置文件路径，方便测试
	prodConfig = "./conf/prod.toml" //配置文件路径，方便测试
	Config     *appCfg              //运行配置实体
	APPNAME    string
	APPVERSION string
	ISTEST     bool
	DSN        string
	RedisAddr  string
	QueueAddr  string
	MongoAddr  string

	ObtenationIterations int = 3

	node, _ = snowflake.NewNode(1)

	IP   string = "192.168.3.235"
	PORT int    = 9080
)

func GetNode() *snowflake.Node {
	return node
}

func GetDsn() string {
	return DSN
}

func GetRedisAddr() string {
	return RedisAddr
}

func GetQueueAddr() string {
	return QueueAddr
}

func GetConfig() *appCfg {
	return Config
}

func GetMongoAddr() string {
	return MongoAddr
}

func init() {
	color.Red("初始化配置")
	var err error
	ISTEST, err = utils.AToBool(os.Getenv("Debug"))
	if err != nil {
		ISTEST = true
	}
	if ISTEST {
		color.Red("当前为测试环境")
		Config, err = initCfg(defConfig)
		if err != nil {
			panic(err)
		}
	} else {
		color.Red("当前为生产环境")
		Config, err = initCfg(prodConfig)
		if err != nil {
			panic(err)
		}
	}
	color.Green("APPNAME=%s", APPNAME)
	APPVERSION = Config.Global.Version
	color.Green("APPVERSION=%s", APPVERSION)
	color.Green("Mysql=%s", Config.Mysql)
}

func initCfg(fn string) (*appCfg, error) {
	app := &appCfg{}
	_, err := toml.DecodeFile(fn, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
