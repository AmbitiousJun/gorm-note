package main

import (
	"gorm-learn/db"
	"log"
)

type Result struct {
	Name string
	Age  int
}

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var result Result
	var results = make([]*Result, 0)

	// 1 查询单条记录
	res := d.Table("users").Select("name", "age").Where("id = ?", 9).Scan(&result)
	if res.Error != nil {
		log.Fatal("查询用户信息出错", res.Error)
	}
	log.Println("查询到 id 为 9 的用户信息：", result)

	// 2 查询多条记录
	res = d.Raw("select name, age from `users` where `deleted_at` is not null").Scan(&results)
	if res.Error != nil {
		log.Fatal("查询用户信息出错", res.Error)
	}
	log.Println("查询所有的用户信息：")
	for _, r := range results {
		log.Println(r)
	}
}
