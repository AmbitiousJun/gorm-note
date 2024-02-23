// 使用 struct, map 构造查询条件
package main

import "gorm-learn/db"

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var findUsers []*db.User

	// 1 使用 struct 构造条件
	d.Where(&db.User{Name: "John", Age: 19}).Find(&findUsers)

	// 2 使用 map 构造条件
	d.Where(map[string]interface{}{"Name": "John", "Age": 19}).Find(&findUsers)

	// 3 使用 slice 查询主键
	d.Where([]int64{20, 21, 22}).Find(&findUsers)

}
