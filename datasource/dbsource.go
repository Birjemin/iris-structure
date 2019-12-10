package datasource

import (
	"database/sql"
	"fmt"
	"github.com/birjemin/iris-structure/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

// redis状态
func StatsDB() sql.DBStats {
	return db.DB().Stats()
}

// 关闭db
func CloseDb() error {
	return db.DB().Close()
}

func init() {
	path := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		conf.Sysconfig.DBUserName,
		conf.Sysconfig.DBPassword,
		conf.Sysconfig.DBIp,
		conf.Sysconfig.DBPort,
		conf.Sysconfig.DBName,
	)
	var err error
	db, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(conf.Sysconfig.DBMaxLifetime * time.Second)
	db.DB().SetMaxIdleConns(conf.Sysconfig.DBMaxIdle) // 设置最大闲置个数
	db.DB().SetMaxOpenConns(conf.Sysconfig.DBMaxOpen) // 最大打开的连接数
	db.SingularTable(true)                            // 表生成结尾不带s
	// 是否启用Logger，显示详细日志
	db.LogMode(conf.Sysconfig.DBDebug)
	// 先注释掉
	//if !db.HasTable(&models.Book{}) { //db.Set 设置一些额外的表属性
	//	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&models.Book{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
}
