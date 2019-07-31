package models

import (
	"fmt"

	. "github.com/jou66jou/go-spider2/logs"
	"github.com/jou66jou/go-spider2/models/mysql"
	"github.com/jou66jou/go-spider2/models/redis"
)

// InitDb init db
func InitDb() error {
	err := mysql.InitMysql()
	if err != nil {
		Log.Errorf("init mysql err: %v", err)
		return fmt.Errorf("init mysql err: %v", err)
	}

	err = redis.InitRedis()
	if err != nil {
		Log.Errorf("init redis err: %v", err)
		return fmt.Errorf("init redis err: %v", err)
	}

	return nil
}

// Close close db
func Close() error {
	mysql.Close()
	redis.Close()
	return nil
}
