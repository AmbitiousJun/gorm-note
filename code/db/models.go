// 定义练习时的通用 Model
package db

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Location struct {
	X, Y int
}

// Scan 实现 sql.Scanner 接口, 用于将数据库数据转换为自定义结构
func (l *Location) Scan(v interface{}) error {
	var bytes []byte
	var ok bool
	if bytes, ok = v.([]byte); !ok {
		return errors.New("数据库返回错误的格式")
	}
	reg := `POINT\((\d+) (\d+)\)`
	exp := regexp.MustCompile(reg)
	results := exp.FindStringSubmatch(string(bytes))
	if len(results) != 3 {
		return errors.New("解析数据库返回结果失败")
	}
	l.X, _ = strconv.Atoi(results[1])
	l.Y, _ = strconv.Atoi(results[2])
	return nil
}

// GormDataType 实现 schema.GormDataTypeInterface 接口, 指定结构体匹配的数据库类型
func (l *Location) GormDataType() string {
	return "geometry"
}

// GormValue 实现 gorm.Valuer 接口, 指定结构体数据应该怎么储存到数据库中
func (l *Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", l.X, l.Y)},
	}
}

func (l *Location) String() string {
	return fmt.Sprintf(`Location{X: %v, Y: %v}`, l.X, l.Y)
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func (cc *CreditCard) String() string {
	return fmt.Sprintf(`CreditCard{ID: %v, Number: %v, UserId: %v}`, cc.ID, cc.Number, cc.UserID)
}

type User struct {
	gorm.Model
	Name       string
	Age        int
	Birthday   time.Time
	Location   *Location
	CreditCard *CreditCard
}

func (u *User) String() string {
	return fmt.Sprintf(`
		User{
			ID: %v,
			CreatedAt: %v,
			UpdatedAt: %v,
			DeletedAt: %v,
			Name: %v,
			Age: %v,
			Birthday: %v,
			Location: %v,
			CreditCard: %v
		}
	`, u.ID, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.Name, u.Age, u.Birthday, u.Location, u.CreditCard)
}
