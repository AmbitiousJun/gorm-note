package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 1 初始化数据库
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm-learn?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	// 2 自动创建数据库
	db.AutoMigrate(&Product{})

	// 3 创建一条记录
	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	// 4 根据主键查询数据
	db.First(&product, 1)
	log.Printf("根据主键 %d 查询到记录: %v\n", 1, product)

	// 5 根据条件查询数据
	db.First(&product, "code = ?", "D42")
	log.Printf("根据条件 %s 查询到记录: %v\n", "code = D42", product)

	// 6 将当前记录的价格修改为 200
	db.Model(&product).Update("Price", 200)
	db.First(&product, 1)
	log.Printf("将 Price 修改为 %d: %v\n", 200, product)

	// 7 使用 struct 一次性修改多个字段
	db.Model(&product).Updates(Product{Price: 300, Code: "F42"})
	db.First(&product, 1)
	log.Printf("使用 struct 一次性修改多个字段: %v\n", product)

	// 8 使用 map 一次性修改多个字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 400, "Code": "G42"})
	db.First(&product, 1)
	log.Printf("使用 map 一次性修改多个字段: %v\n", product)

	// 9 根据主键删除记录
	db.Delete(&product, 1)
	db.First(&product, 1)
	log.Printf("删除记录后的查询结果: %v\n", product)
}
