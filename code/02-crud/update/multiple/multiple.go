// 更新多个字段
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var findUser = new(db.User)
	d.First(findUser, 27)
	log.Println("更新多个字段 => 用户原始信息: ", findUser)

	// 1 使用 struct 更新
	result := d.Model(findUser).Updates(db.User{Name: "王五", Age: 22})
	log.Println("struct 更新多个字段 => 错误信息: ", result.Error)
	log.Println("struct 更新多个字段 => 影响行数: ", result.RowsAffected)
	d.First(findUser, 27)
	log.Println("struct 更新多个字段 => 用户信息重查: ", findUser)

	// 2 使用 map 更新
	result = d.Model(findUser).Updates(map[string]interface{}{"name": "赵六", "age": 30})
	log.Println("map 更新多个字段 => 错误信息: ", result.Error)
	log.Println("map 更新多个字段 => 影响行数: ", result.RowsAffected)
	d.First(findUser, 27)
	log.Println("map 更新多个字段 => 用户信息重查: ", findUser)
}
