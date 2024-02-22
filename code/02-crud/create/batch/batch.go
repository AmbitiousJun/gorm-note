// 批量插入
package main

import (
	"gorm-learn/db"
	"log"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	users := []db.User{
		{Name: "ZhangSan", Age: 18},
		{Name: "LiSi", Age: 19},
		{Name: "WangWu", Age: 20},
		{Name: "ZhaoLiu", Age: 21},
		{Name: "TianQi", Age: 22},
	}

	// 一次插入 2 条数据
	result := d.Omit("Birthday").CreateInBatches(users, 2)
	log.Println("批量插入数据 => 错误信息: ", result.Error)
	log.Println("批量插入数据 => 影响行数: ", result.RowsAffected)
}
