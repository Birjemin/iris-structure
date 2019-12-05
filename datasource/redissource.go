package datasource

import (
	"fmt"
	"github.com/birjemin/iris-structure/conf"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// redis状态
func StatsRedis() redis.PoolStats {
	return redisPool.Stats()
}

// 关闭redis
func CloseRedis() error {
	if redisPool != nil {
		return redisPool.Close()
	}
	return nil
}

// 初始化redis
func InitRedis() {
	redisPool = &redis.Pool{
		MaxIdle:     conf.Sysconfig.RedisMaxIdle,
		MaxActive:   conf.Sysconfig.RedisMaxOpen,
		IdleTimeout: conf.Sysconfig.RedisMaxLifetime * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", conf.Sysconfig.RedisIp, conf.Sysconfig.RedisPort),
				redis.DialDatabase(conf.Sysconfig.RedisDB),
				redis.DialPassword(conf.Sysconfig.RedisPassword),
			)
		},
	}
	conn := GetRedis()
	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		panic("redis connect failed.")
	}
}
