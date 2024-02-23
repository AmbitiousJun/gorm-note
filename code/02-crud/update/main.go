// 基础的更新方法
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 更新所有字段
	var user = new(db.User)
	d.First(user, 27)
	log.Println("更新所有字段 => 查询出的用户原信息: ", *user)
	user.Name = "李四"
	user.Age = 99
	result := d.Save(user)
	log.Println("更新所有字段 => 错误信息: ", result.Error)
	log.Println("更新所有字段 => 影响行数: ", result.RowsAffected)
	d.First(user, 27)
	log.Println("更新所有字段 => 修改后重新查询用户信息: ", *user)
}
