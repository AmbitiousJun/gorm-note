package main

import (
	"gorm-learn/db"
	"log"
)

// 使用 User 结构体的 "子集" 结构体进行查询，GORM 会智能开启部分字段查询
type APIUser struct {
	ID   uint
	Name string
}

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var users = make([]*APIUser, 0)
	res := d.Model(&db.User{}).Limit(10).Find(&users)

	if res.Error != nil {
		log.Fatal("查询用户失败", res.Error)
	}
	log.Println("查询到的用户数据")
	for _, user := range users {
		log.Println(user)
	}
}
