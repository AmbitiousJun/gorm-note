// 使用 SQL 表达式进行更新
package main

import (
	"gorm-learn/db"
	"log"

	"gorm.io/gorm"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var findUser = new(db.User)
	d.First(findUser, 27)
	log.Println("sql 表达式更新 => 原始用户信息: ", findUser)

	result := d.Model(findUser).Update("age", gorm.Expr("age * ? + ?", 2, 100))
	log.Println("sql 表达式更新 => 错误信息: ", result.Error)
	log.Println("sql 表达式更新 => 影响行数: ", result.RowsAffected)
	d.First(findUser, 27)
	log.Println("sql 表达式更新 => 用户信息重查: ", findUser)
}
