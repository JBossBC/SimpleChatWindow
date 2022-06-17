package DataSource

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

type RedisConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

var redisConnection *redis.Conn
var redisPool *redis.Pool

func RedisInit() {
	redisConfig := GetRedisConfig()
	redisStr := strings.Builder{}
	redisStr.WriteString(redisConfig.Hostname)
	redisStr.WriteString(":")
	redisStr.WriteString(redisConfig.Port)
	var err error
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", redisStr.String())
			return dial, err
		},
		TestOnBorrow:    nil,
		MaxIdle:         0,
		MaxActive:       255,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	if err != nil {
		panic("redis初始化错误")
	}
	redisPool = &pool
}

func GetRedisConnection() redis.Conn {
	return redisPool.Get()
}
