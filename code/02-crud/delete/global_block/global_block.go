// 默认全局删除阻塞
package main

import (
	"gorm-learn/db"
	"log"

	"gorm.io/gorm"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 调用删除方法时没有指定 where 条件会直接抛出异常
	result := d.Delete(&db.User{})
	log.Println("没有指定 where 条件 => 错误信息: ", result.Error)
	log.Println("没有指定 where 条件 => 影响行数: ", result.RowsAffected)

	// 绕过方式 1：指定一个永远为真的条件
	result = d.Where("1 = 1").Delete(&db.User{})
	log.Println("绕过方式 1 => 错误信息: ", result.Error)
	log.Println("绕过方式 1 => 影响行数: ", result.RowsAffected)

	// 绕过方式 2：执行原始 SQL
	result = d.Exec("delete from users")
	log.Println("绕过方式 2 => 错误信息: ", result.Error)
	log.Println("绕过方式 2 => 影响行数: ", result.RowsAffected)

	// 绕过方式 3：开启 AllowGlobalUpdate
	result = d.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&db.User{})
	log.Println("绕过方式 3 => 错误信息: ", result.Error)
	log.Println("绕过方式 3 => 影响行数: ", result.RowsAffected)
}
