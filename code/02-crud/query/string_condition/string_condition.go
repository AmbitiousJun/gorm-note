// 字符串形式的查询条件构造
package main

import (
	"gorm-learn/db"
	"time"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var findUser = new(db.User)
	var findUsers []db.User

	// Equals
	d.Where("name = ?", "jinzhu").First(findUser)

	// Not Equals
	d.Where("name <> ?", "jinzhu").Find(&findUsers)

	// In
	d.Where("name IN ?", []string{"jinzhu", "jinzhu_2"}).Find(&findUsers)

	// Like
	d.Where("name LIKE ?", "%jin%").Find(&findUsers)

	// And
	d.Where("name = ? AND age >= ?", "jinzhu", 22).Find(&findUsers)

	lastWeek, today := time.Now().AddDate(0, 0, -7), time.Now()
	// Time
	d.Where("updated_ad > ?", lastWeek).Find(&findUsers)

	// Between
	d.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&findUsers)
}
