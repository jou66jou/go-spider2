package redis

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jou66jou/go-spider2/conf"
)

type RedisPool struct {
	masterPool *redis.Pool
}

var (
	redisPool *RedisPool
)

// InitRedis init redis
func InitRedis() error {
	masterRedisPool, err := InitMasterRedis()
	if err != nil {
		return errors.New("init master redis err: " + err.Error())
	}

	redisPool = &RedisPool{
		masterPool: masterRedisPool,
	}

	return nil
}

// InitMasterRedis init master redis
func InitMasterRedis() (*redis.Pool, error) {
	proto := "tcp"

	pool := &redis.Pool{
		MaxIdle:     conf.AppConf.RedisConf.RedisMaxIdle,
		MaxActive:   conf.AppConf.RedisConf.RedisMaxActive,
		IdleTimeout: time.Duration(conf.AppConf.RedisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(proto, conf.AppConf.RedisConf.RedisAddr)
			if err != nil {
				tmpStr := fmt.Sprintf("redis proto=%v addr=%v dial err: %v", proto, conf.AppConf.RedisConf.RedisAddr, err)
				return nil, errors.New(tmpStr)
			}

			if _, err = conn.Do("AUTH", conf.AppConf.RedisConf.RedisAuth); err != nil {
				conn.Close()
				tmpStr := fmt.Sprintf("redis set addr=%v auth=%v is err: %v", conf.AppConf.RedisConf.RedisAddr, conf.AppConf.RedisConf.RedisAuth, err)
				return nil, errors.New(tmpStr)
			}
			return conn, nil
		},
	}

	c := pool.Get()
	defer c.Close()
	c.Do("ERR", io.EOF)
	if c.Err() != nil {
		tmpStr := fmt.Sprintf("redis do err: %v", c.Err())
		return nil, errors.New(tmpStr)
	}

	return pool, nil
}

// SetNotExKey 不過期 key
func SetNotExKey(key, value string) error {
	c := redisPool.masterPool.Get()
	defer c.Close()

	_, err := c.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// SetNotExKey 設置過期 key
func SetExKey(key, value string, exTime int) error {
	c := redisPool.masterPool.Get()
	defer c.Close()

	_, err := c.Do("set", key, value, "ex", exTime)
	if err != nil {
		return err
	}

	return nil
}

// GetDbKey 取得 value
func GetDbKey(key string) (string, error) {
	c := redisPool.masterPool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return "", errors.New("key not exist")
		}
		return "", err
	}

	return value, nil
}

// GetDbKey 查詢key 存在 bool
func GetDbKeyExist(key string) (bool, error) {
	c := redisPool.masterPool.Get()
	defer c.Close()

	return redis.Bool(c.Do("EXISTS", key))
}

// DelDbKey 刪除 key
func DelDbKey(key string) error {
	c := redisPool.masterPool.Get()
	defer c.Close()

	_, err := c.Do("del", key)
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	redisPool.masterPool.Close()
}
