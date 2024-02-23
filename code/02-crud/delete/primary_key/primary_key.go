// 根据主键进行删除
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 1 传递 int 类型的主键值进行删除
	result := d.Delete(&db.User{}, 1)
	log.Println("传递 int 类型的主键值进行删除 => 错误信息: ", result.Error)
	log.Println("传递 int 类型的主键值进行删除 => 影响行数: ", result.RowsAffected)

	// 2 传递 string 类型的主键值进行删除
	result = d.Delete(&db.User{}, "2")
	log.Println("传递 string 类型的主键值进行删除 => 错误信息: ", result.Error)
	log.Println("传递 string 类型的主键值进行删除 => 影响行数: ", result.RowsAffected)

	// 3 传递 slice 类型的主键值进行批量删除
	result = d.Delete(&db.User{}, []int{3, 4, 5})
	log.Println("传递 slice 类型的主键值进行删除 => 错误信息: ", result.Error)
	log.Println("传递 slice 类型的主键值进行删除 => 影响行数: ", result.RowsAffected)
}
