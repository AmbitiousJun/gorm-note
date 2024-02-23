// 插入或更新时冲突处理
package main

import (
	"gorm-learn/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	d := db.DB()
	d.AutoMigrate(&db.User{})

	user := db.User{Name: "Sarra", Age: 30}

	// 1 忽略冲突
	d.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// 2 当字段 id 冲突时，修改指定字段为默认值
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL
	d.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
	}).Create(&user)

	// 3 当字段 id 冲突时，使用 SQL 生成指定字段值
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `count`=GREATEST(count, VALUES(count));
	d.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
	}).Create(&user)

	// 4 当字段 id 冲突时，将指定字段更新为新值
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age); MySQL
	d.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(&user)

	// 5 当字段 id 冲突时，将除了主键之外的全部字段进行更新
	d.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user)
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age", ...;
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age), ...; MySQL
}
