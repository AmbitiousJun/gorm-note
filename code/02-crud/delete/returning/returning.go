// 删除记录的时候返回被删除的数据
package main

import (
	"gorm-learn/db"
	"log"

	"gorm.io/gorm/clause"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	var users []db.User
	result := d.Clauses(clause.Returning{
		Columns: []clause.Column{{Name: "age"}, {Name: "name"}},
	}).Delete(&users, "age in ?", []int{18, 20})
	log.Println("删除时返回数据 => 错误信息: ", result.Error)
	log.Println("删除时返回数据 => 影响行数: ", result.RowsAffected)
	log.Println("删除时返回数据 => users: ", users)
}
