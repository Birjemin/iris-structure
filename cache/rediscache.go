package cache

import (
	"errors"
	"github.com/birjemin/iris-structure/datasource"
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/golog"
)

//设置缓存
func Set(key, val string, ttl int) error {
	conn := datasource.GetRedis()
	defer conn.Close()
	var r string
	var err error
	if ttl <= 0 {
		r, err = redis.String(conn.Do("SET", key, val))
	} else {
		r, err = redis.String(conn.Do("SET", key, val, "EX", ttl))
	}

	if err != nil {
		golog.Errorf("[rediscache]method:Set,err: %s", err)
		return err
	}

	if r != "OK" {
		return errors.New("NOT OK")
	}

	return nil
}

//获取缓存
func Get(key string) (string, error) {
	conn := datasource.GetRedis()
	defer conn.Close()
	r, err := redis.String(conn.Do("GET", key))
	if err != nil {
		golog.Errorf("[rediscache]method:Get,err: %s", err)
		return "", err
	}

	return r, nil
}

func LPop(queue string) (string, error) {
	conn := datasource.GetRedis()
	defer conn.Close()
	r, err := redis.String(conn.Do("lpop", queue))
	if err != nil {
		return "", err
	}
	return r, nil
}

