package conf

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

var Sysconfig = &sysconfig{}

func init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = json.Unmarshal(b, Sysconfig)
	if err != nil {
		panic(err)
	}
}

type sysconfig struct {
	LoggerLevel string `json:"LoggerLevel"`
	// mysql
	DBDebug       bool          `json:"DBDebug"`
	DBUserName    string        `json:"DBUserName"`
	DBPassword    string        `json:"DBPassword"`
	DBIp          string        `json:"DBIp"`
	DBPort        int           `json:"DBPort"`
	DBName        string        `json:"DBName"`
	DBMaxIdle     int           `json:"DBMaxIdle"`     // 最大的空闲连接数
	DBMaxOpen     int           `json:"DBMaxOpen"`     // 最大连接数
	DBMaxLifetime time.Duration `json:"DBMaxLifetime"` // 生命周期
	// redis
	RedisIp          string        `json:"RedisIp"`
	RedisPort        int           `json:"RedisPort"`
	RedisDB          int           `json:"RedisDB"`
	RedisPassword    string        `json:"RedisPassword"`
	RedisMaxIdle     int           `json:"RedisMaxIdle"`     // 最大的空闲连接数
	RedisMaxOpen     int           `json:"RedisMaxOpen"`     // 最大连接数
	RedisMaxLifetime time.Duration `json:"RedisMaxLifetime"` // 生命周期
	ConsumerNum      int           `json:"ConsumerNum"`
}
