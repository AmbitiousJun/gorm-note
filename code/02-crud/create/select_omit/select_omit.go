// 指定字段新增
package main

import (
	"gorm-learn/db"
	"log"
	"time"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var findUser = new(db.User)

	// 1 指定插入特定的字段
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	user := db.User{Name: "John", Age: 19, Birthday: time.Now()}
	result := d.Select("Name", "Age", "CreatedAt").Create(&user)
	log.Println("指定插入特定字段 => 错误信息: ", result.Error)
	log.Println("指定插入特定字段 => 影响行数: ", result.RowsAffected)
	d.First(findUser, user.ID)
	log.Println("指定插入特定字段 => 新增结果: ", findUser)

	// 2 指定不插入特定字段
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
	user.ID = 0
	result = d.Omit("Name", "Age", "CreatedAt").Create(&user)
	log.Println("指定不插入特定字段 => 错误信息: ", result.Error)
	log.Println("指定不插入特定字段 => 影响行数: ", result.RowsAffected)
	findUser = new(db.User)
	d.First(findUser, user.ID)
	log.Println("指定不插入特定字段 => 新增结果: ", findUser)
}
