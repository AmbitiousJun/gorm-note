// 使用 map 创建记录
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 1 创建单条记录
	result := d.Model(&db.User{}).Create(map[string]interface{}{
		"Name": "Johnny",
		"Age":  18,
	})
	log.Println("map 创建单条记录 => 错误信息: ", result.Error)
	log.Println("map 创建单条记录 => 影响行数: ", result.RowsAffected)

	// 2 创建多条记录
	result = d.Model(&db.User{}).Create([]map[string]interface{}{
		{"Name": "aaa", "Age": 20},
		{"Name": "bbb", "Age": 21},
	})
	log.Println("map 创建多条记录 => 错误信息: ", result.Error)
	log.Println("map 创建多条记录 => 影响行数: ", result.RowsAffected)
}
