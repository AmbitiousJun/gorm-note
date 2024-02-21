# GORM 学习笔记

> 文档地址：https://gorm.io/docs/
> 
> 仓库地址：https://github.com/AmbitiousJun/gorm-note

## 安装 & 快速入门

### 安装

需要安装两个库，分别是 gorm 和 mysql 数据库的驱动包

在终端中按顺序执行一下命令进行安装：

```shell
# gorm 库
go get -u gorm.io/gorm
# mysql 驱动包
go get -u gorm.io/driver/mysql
```

在编写程序之前，需要先创建一个 mysql 数据库，这里命名为 `gorm-learn`

![](assets/2024-02-21-10-32-25-image.png)

无需创建表，gorm 提供了 API 可以方便地根据实体类生成相对应的表结构

接下来就是编写程序连接 mysql 进行增删改查操作了，首先建好一个 main.go 文件，写好一个空的 main 函数

### 连接

首先是连接数据库，使用 `gorm.Open()` 函数打开连接，接收两个参数，第一个是**数据库连接驱动**，第二个是**定制化配置**

要连接 mysql 数据库，就使用 mysql 驱动，即调用 `mysql.Open()` 函数，传入数据库连接 url 即可

不需要定制化配置，就传递一个空结构体

连接成功之后，可以获取到一个 `*gorm.DB` 类型的对象，这是用于操作数据库的核心对象

```go
// 1 初始化数据库
db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm-learn?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
if err != nil {
    log.Fatalln("failed to connect database")
}
```

### 自动生成表结构

使用 gorm 不需要手动建表，只需要在代码中维护好结构体的属性即可，下面定义一个 `Product` 结构体：

```go
type Product struct {
    gorm.Model
    Code  string
    Price uint
}
```

gorm 提供了一个基础的结构体 `gorm.Model` ，包含了主键 (ID)、创建时间 (CreateAt)、更新时间 (UpdateAt)、删除时间 (DeleteAt) 等信息，可以直接将其**组合**到自定义的结构体下

调用 `db.AutoMigrate` 方法就可以很方便地创建表结构，表名是小写的结构体名称加上 s，例如现在这个结构体会自动生成 `products` 表

```go
// 2 自动创建数据库
db.AutoMigrate(&Product{})
```

### 新增

调用方法：`db.Create()`

```go
db.Create(&Product{Code: "D42", Price: 100})
```

### 查询

调用方法：`db.First()`

传入单个参数时，gorm 会去自动匹配主键字段进行查询

传入两个参数时，gorm 会将第一个参数作为预编译 sql，第二个参数作为 sql 中的数据值进行查询

```go
var product Product
// 4 根据主键查询数据
db.First(&product, 1)
log.Printf("根据主键 %d 查询到记录: %v\n", 1, product)

// 5 根据条件查询数据
db.First(&product, "code = ?", "D42")
log.Printf("根据条件 %s 查询到记录: %v\n", "code = D42", product)
```

### 修改

调用方法：`db.Model().Update()` | `db.Model().Updates`

首先通过 `Model` 方法指定要更新哪条记录

接着链式调用 `Update[s]` 方法进行更新

`Update` 方法用于更新单个字段，接收两个参数，分别是字段名和字段值 

`Updates` 方法用于更新多个字段，接收一个参数，只能是 **struct** 和 **map** 类型

```go
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
```

### 删除

调用方法：`db.Delete()`

传入要删除记录的主键值即可删除对应的记录

> gorm 会自动检测表中是否有 `DeleteAt` 属性，有的话在删除的时候就会采用逻辑删除模式，给 `DeleteAt` 属性设置值，而不实际删除记录
> 
> 在查询时，同样地会去检查当前字段是否为空，不为空则认为记录不存在
> 
> 由于 Product 结构体组合了 gorm 提供的 Model 结构，Model 结构中包含了 `DeleteAt` 属性，所以会自动采用逻辑删除模式

```go
// 9 根据主键删除记录
db.Delete(&product, 1)
```

### 完整代码

```go
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
```

### 运行结果

![](assets/2024-02-21-11-29-42-image.png)

![](assets/2024-02-21-11-30-10-image.png)
