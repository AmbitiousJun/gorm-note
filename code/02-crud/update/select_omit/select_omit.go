// 使用 Select/Omit 部分更新数据
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
	log.Println("原始的用户信息: ", findUser)

	// 1 select 部分更新
	result := d.Model(findUser).Select("Name").Updates(db.User{Name: "啊啊啊", Age: 50})
	log.Println("select 部分更新 => 错误信息: ", result.Error)
	log.Println("select 部分更新 => 影响行数: ", result.RowsAffected)
	d.First(findUser, 27)
	log.Println("select 部分更新 => 用户信息重查: ", findUser)

	// 2 omit 忽略更新
	result = d.Model(findUser).Omit("Name").Updates(db.User{Name: "嗯嗯嗯", Age: 24})
	log.Println("omit 忽略更新 => 错误信息: ", result.Error)
	log.Println("omit 忽略更新 => 影响行数: ", result.RowsAffected)
	d.First(findUser, 27)
	log.Println("omit 忽略更新 => 用户信息重查: ", findUser)

}
