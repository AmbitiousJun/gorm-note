package main

import (
	"gorm-learn/db"
	"log"
	"math"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var users = make([]*db.User, 0)
	var users1 = make([]*db.User, 0)

	// 1 查询所有用户, 限制最多查询 3 条记录
	res := d.Limit(3).Find(&users)
	if res.Error != nil {
		log.Fatal("01 => 查询失败", res.Error)
	}
	for _, user := range users {
		log.Println(user)
	}

	// 2 链式调用, 先查询所有用户, 限制最多查询 10 条记录
	//   再取消限制, 然后重新查询
	res = d.Limit(10).Find(&users).Limit(-1).Find(&users1)
	if res.Error != nil {
		log.Fatal("02 => 查询失败", res.Error)
	}
	log.Printf("02 => 查询所有用户, 限制最多查询 10 条记录, 第一次查询到了 %d 条数据\n", len(users))
	log.Printf("02 => 第二次查询到了 %d 条数据\n", len(users1))

	// 3 从第 4 条记录开始, 查询所有用户
	res = d.Limit(math.MaxInt32).Offset(3).Find(&users)
	if res.Error != nil {
		log.Fatal("03 => 查询失败", res.Error)
	}
	log.Printf("03 => 从第 4 条记录开始查询到了 %d 条记录, 首个用户 id 为 %d\n", len(users), users[0].ID)

	// 4 从第 6 条记录开始查询所有用户, 最多查询 10 条记录
	res = d.Offset(5).Limit(10).Find(&users)
	if res.Error != nil {
		log.Fatal("04 => 查询失败", res.Error)
	}
	log.Printf("04 => 从第 6 条记录开始, 最多查询 10 条记录, 查询到了 %d 条记录, 首个用户 id 为 %d\n", len(users), users[0].ID)

	// 5 链式调用, 先查询所有用户, 从第 11 条记录开始查询
	//   再取消限制, 然后重新查询
	res = d.Limit(math.MaxInt32).Offset(10).Find(&users).Offset(-1).Find(&users1)
	if res.Error != nil {
		log.Fatal("05 => 查询失败", res.Error)
	}
	log.Printf("05 => 从第 11 条记录开始, 查询所有用户, 查询到了 %d 条记录, 再取消限制, 查询到了 %d 条记录\n", len(users), len(users1))
}
