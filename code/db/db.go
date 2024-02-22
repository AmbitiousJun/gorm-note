// 统一的 database 管理
package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var initOnce sync.Once

func DB() *gorm.DB {
	initOnce.Do(func() {
		d, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm-learn?charset=utf8mb4&parseTime=True&loc=Local"))
		if err != nil {
			log.Fatal("failed to connect database")
		}
		db = d
	})
	return db
}
