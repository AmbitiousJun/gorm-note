// 批量删除
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 传递一个不包含主键的实体，自动执行批量删除
	result := d.Delete(&db.User{}, "name like ?", "%Haha%")
	log.Println("批量删除 => 错误信息: ", result.Error)
	log.Println("批量删除 => 影响行数: ", result.RowsAffected)
}
