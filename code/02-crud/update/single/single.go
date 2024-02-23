// 更新单个字段
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 使用自定义的 where 条件进行更新
	result := d.Model(&db.User{}).Where("name = ?", "李四").Update("age", 88)
	log.Println("更新单个字段 => 错误信息: ", result.Error)
	log.Println("更新单个字段 => 影响行数: ", result.RowsAffected)
	var findUser = new(db.User)
	d.First(findUser, "name = ?", "李四")
	log.Println("更新单个字段 => 重查数据: ", *findUser)
}
