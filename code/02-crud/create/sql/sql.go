// 使用自定义的 SQL 表达式进行查询
package main

import (
	"gorm-learn/db"
	"log"
	"time"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	// 创建一条带有 Location 属性的记录
	user := db.User{
		Name:     "Haha",
		Age:      30,
		Birthday: time.Now(),
		Location: &db.Location{X: 33, Y: 44},
	}
	result := d.Create(&user)
	log.Println("通过 SQL 表达式创建 => 错误信息: ", result.Error)
	log.Println("通过 SQL 表达式创建 => 影响行数: ", result.RowsAffected)

	// 查询
	var findUser = new(db.User)
	d.Raw("select id, created_at, updated_at, deleted_at, name, age, birthday, ST_AsText(location) as location from users where id = ?", user.ID).Scan(&findUser)
	log.Println("通过 SQL 表达式创建 => 重查结果: ", findUser)
}
