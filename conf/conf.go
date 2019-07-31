package conf

import (
	"fmt"

	ini "gopkg.in/ini.v1"
)

var (
	AppConf Config
)

type Config struct {
	DbConf
	RedisConf
	NsqConf
	LogConf
}

// mysql db config
type DbConf struct {
	DbName        string `ini:"DbName"`
	DbHost        string `ini:"DbHost"`
	DbPort        string `ini:"DbPort"`
	DbUser        string `ini:"DbUser"`
	DbPassword    string `ini:"DbPassword"`
	DbLogEnable   bool   `ini:"DbLogEnable"`
	DbMaxConnect  int    `ini:"DbMaxConnect"`
	DbIdleConnect int    `ini:"DbIdleConnect"`
}

// redis db config
type RedisConf struct {
	RedisAddr        string `ini:"RedisAddr"`
	RedisAuth        string `ini:"RedisAuth"`
	RedisMaxIdle     int    `ini:"RedisMaxIdle"`
	RedisMaxActive   int    `ini:"RedisMaxActive"`
	RedisIdleTimeout int    `ini:"RedisIdleTimeout"`
}

// nsq
type NsqConf struct {
	ProducerAddr  string `ini:"ProducerAddr"`
	ProducerTopic string `ini:"ProducerTopic"`
}

// Log config
type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
}

func InitConf(confPath string) error {
	AppConf = Config{}
	if err := ini.MapTo(&AppConf, confPath); err != nil {
		return err
	}
	return nil
}

func GetDbConStr() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConf.DbUser, AppConf.DbPassword, AppConf.DbHost, AppConf.DbPort, AppConf.DbName,
	)
}
