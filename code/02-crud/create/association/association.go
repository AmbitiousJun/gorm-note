// 关联创建
package main

import (
	"gorm-learn/db"
	"log"
	"time"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})
	d.AutoMigrate(&db.CreditCard{})

	user := db.User{
		Name:     "张三",
		Age:      30,
		Birthday: time.Now(),
		CreditCard: &db.CreditCard{
			Number: "188282374893789378",
		},
	}

	result := d.Create(&user)
	log.Println("关联创建 => 错误信息: ", result.Error)
	log.Println("关联创建 => 影响行数: ", result.RowsAffected)

	var findUser = new(db.User)
	d.First(findUser, user.ID)
	log.Println("关联创建 => 查询结果: ", findUser)
}
