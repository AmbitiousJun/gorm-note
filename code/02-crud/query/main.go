// 基本的查询方法
package main

import (
	"errors"
	"gorm-learn/db"
	"log"

	"gorm.io/gorm"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 1 查询 1 条记录, 根据主键升序排序
	var findUser = new(db.User)
	result := d.First(findUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal("查询不到记录: ", result.Error)
	} else {
		log.Println("First 查询 => 用户信息: ", findUser)
	}

	// 2 查询 1 条记录, 不排序
	findUser = new(db.User)
	result = d.Take(findUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal("查询不到记录: ", result.Error)
	} else {
		log.Println("Take 查询 => 用户信息: ", findUser)
	}

	// 3 查询 1 条记录, 根据主键降序排序
	findUser = new(db.User)
	result = d.Last(findUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal("查询不到记录: ", result.Error)
	} else {
		log.Println("Last 查询 => 用户信息: ", findUser)
	}

}
