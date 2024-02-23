// 查询所有记录
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var users []db.User
	result := d.Find(&users)
	log.Println("查询所有记录 => 错误信息: ", result.Error)
	log.Println("查询所有记录 => 用户数: ", len(users))
	for _, user := range users {
		log.Println(user)
	}
}
