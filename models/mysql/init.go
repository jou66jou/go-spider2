package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jou66jou/go-spider2/common"
	"github.com/jou66jou/go-spider2/conf"
	. "github.com/jou66jou/go-spider2/logs"
)

type Database struct {
	Dbs *gorm.DB
}

type newTable struct { //新建各個彩的表
	DataResult
	Isuutime_stamp int64  `gorm:"column:issuetime_stamp"`
	lotteryCode    string `gorm:"-"`
}

func (nt *newTable) TableName() string { //自訂newTable的表名
	return common.TablePrefix + nt.lotteryCode
}

var db *Database

func InitMysql() error {
	Log.Info("init mysql db..............")

	platDb, err := initPlatDb()
	if err != nil {
		return fmt.Errorf("init mysql db err: %v", err)
	}

	db = &Database{
		Dbs: platDb,
	}

	return nil
}

func initPlatDb() (*gorm.DB, error) {
	dbConf := conf.AppConf.DbConf
	tmpDb, err := gorm.Open("mysql", conf.GetDbConStr())
	if err != nil {
		tmpStr := fmt.Sprintf("Connet the db err, dbHost: %v, dbPort: %v, err: %v", dbConf.DbHost, dbConf.DbPort, err)
		Log.Errorf(tmpStr)
		return nil, fmt.Errorf(tmpStr)
	}

	tmpDb.DB().SetMaxOpenConns(dbConf.DbMaxConnect)
	tmpDb.DB().SetMaxIdleConns(dbConf.DbIdleConnect)

	if err = tmpDb.DB().Ping(); err != nil {
		tmpStr := fmt.Sprintf("Ping the db, dbHost: %v, dbPort: %v, err: %v", dbConf.DbHost, dbConf.DbPort, err)
		Log.Error(tmpStr)
		return nil, fmt.Errorf(tmpStr)
	}

	tmpDb.LogMode(dbConf.DbLogEnable)
	tmpDb.SingularTable(true)

	tables := []interface{}{}

	tmpDb = tmpDb.AutoMigrate(tables...)

	// 自動創建表
	// for _, v := range common.LotteryTable {
	// 	t := &newTable{}
	// 	t.lotteryCode = v
	// 	tmpDb.CreateTable(t)
	// }
	for _, v := range tables {
		if !tmpDb.HasTable(v) {
			Log.Errorf("build table %v failed", v)
			return nil, fmt.Errorf("build table %v failed", v)
		}
	}

	return tmpDb, nil
}

func Close() error {
	db.Dbs.Close()
	return nil
}
