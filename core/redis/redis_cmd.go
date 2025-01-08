package redis

import (
	"AdServerCollector/conf"
	"crypto/tls"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

// Config represents Redis configuration.
type Config struct {
	Addrs        []string      `json:"addrs"`      //集群地址
	Password     string        `json:"password"`   //密码
	KeyPrefix    string        `json:"key_prefix"` //key前缀
	DbIndex      int           `json:"db_index"`   //数据索引
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
}

type RedisCmd interface {
	redis.Cmdable
}

func NewRedisCmd(c *Config) (RedisCmd, error) {
	if len(c.Addrs) == 0 {
		return nil, ErrRedisAddrsEmpty
	}

	if c.UseCluster {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:              c.Addrs,
			ReadOnly:           c.ReadOnly,
			Password:           c.Password,
			DialTimeout:        c.DialTimeout,
			ReadTimeout:        c.ReadTimeout,
			WriteTimeout:       c.WriteTimeout,
			PoolSize:           c.PoolSize,
			PoolTimeout:        c.PoolTimeout,
			IdleTimeout:        c.IdleTimeout,
			IdleCheckFrequency: c.IdleCheckFrequency,
		}), nil
	}
	return redis.NewClient(&redis.Options{
		Addr:               c.Addrs[0],
		Password:           c.Password,
		DB:                 c.DbIndex,
		DialTimeout:        c.DialTimeout,
		ReadTimeout:        c.ReadTimeout,
		WriteTimeout:       c.WriteTimeout,
		PoolSize:           c.PoolSize,
		PoolTimeout:        c.PoolTimeout,
		IdleTimeout:        c.IdleTimeout,
		IdleCheckFrequency: c.IdleCheckFrequency,
		//ReadOnly:           c.ReadOnly,
		TLSConfig: c.TLSConfig,
	}), nil
}

var redisCon RedisCmd

func init() {
	_redisConfig := conf.Config.Redis
	_conf, err := json.Marshal(_redisConfig)
	if err != nil {
		panic(err)
	}
	var redisConf Config
	err = json.Unmarshal(_conf, &redisConf)
	redisCon, err = NewRedisCmd(&redisConf)
}

func GetRedisDb() RedisCmd {
	return redisCon
}
