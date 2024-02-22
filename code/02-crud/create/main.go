// 基本的创建语法
package main

import (
	"fmt"
	"gorm-learn/db"
	"log"
	"strings"
	"time"
)

func main() {
	db.DB().AutoMigrate(&db.User{})

	// 1 创建 user, 获取主键以及操作结果
	user := db.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.DB().Create(&user)
	log.Println("新增用户的 ID: ", user.ID)
	log.Println("新增时的错误: ", result.Error)
	log.Println("新增时的数据库影响行数: ", result.RowsAffected)

	// 2 批量新增 user
	users := []*db.User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	result = db.DB().Create(users)
	ids := []string{}
	for _, user := range users {
		ids = append(ids, fmt.Sprintf("%v", user.ID))
	}
	log.Printf("批量新增的用户 Id: %s", strings.Join(ids, ", "))
	log.Println("批量新增时的错误: ", result.Error)
	log.Println("批量新增时的数据库影响行数: ", result.RowsAffected)
}
