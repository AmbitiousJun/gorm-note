// 基本的删除语法
package main

import (
	"gorm-learn/db"
	"log"
	"time"
)

func main() {
	user := db.User{Name: "ZhangSan", Age: 28, Birthday: time.Now()}

	d := db.DB()
	d.AutoMigrate(&user)
	d.Create(&user)

	// 1 直接传递实体进行删除，自动根据 id 删除
	result := d.Delete(&user)
	log.Println("传递实体进行删除 => 错误信息: ", result.Error)
	log.Println("传递实体进行删除 => 影响行数: ", result.RowsAffected)

	// 2 自定义删除条件 + id
	result = d.Where("name = ?", "ZhangSan").Delete(&user)
	log.Println("自定义删除条件删除 => 错误信息: ", result.Error)
	log.Println("自定义删除条件删除 => 影响行数: ", result.RowsAffected)
}
