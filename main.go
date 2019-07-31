// 架構完整爬蟲微服務

package main

import (
	"flag"
	"fmt"

	"github.com/jou66jou/go-spider2/conf"
	"github.com/jou66jou/go-spider2/logs"
	"github.com/jou66jou/go-spider2/models"
)

var (
	confPath = flag.String("config", "./conf/app.ini", "spider profilePath")
)

func Init() error {
	flag.Parse()
	var err error
	err = conf.InitConf(*confPath)
	if err != nil {
		return fmt.Errorf("init config err: %v", err)
	}

	err = logs.InitLog(conf.AppConf.LogPath, conf.AppConf.LogLevel)
	if err != nil {
		return fmt.Errorf("init log is err: %v", err)
	}
	// 初始化 mysql 和 redis 数据库
	err = models.InitDb()
	if err != nil {
		return fmt.Errorf("init db is err: %v", err)
	}

	return nil
}

func main() {

}
