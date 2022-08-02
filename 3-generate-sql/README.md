## Day3 - 生成sql语句(新增和查询)
###学习反射
https://blog.csdn.net/xhd731568849/article/details/79198048
```
3-generate-sql/
    |--generate/          # 日志
        |--caluse.go    
    |--session/      # 数据库交互
        |--raw.go
    |--orm.go     # 用户交互
    |--go.mod
```

### clause
clause子句生成
- `Set` 方法根据 `Type` 调用对应的 `generator`，生成该子句对应的 SQL 语句。
- `Build` 方法根据传入的 `Type` 的顺序，构造出最终的 SQL 语句。

### 实现Insert
ORM框架中期望Insert的调用方式
```go
s := geeorm.NewEngine("sqlite3", "gee.db").NewSession()
u1 := &User{Name: "Tom", Age: 18}
u2 := &User{Name: "Sam", Age: 25}
s.Insert(u1, u2, ...)
```
完成u1、u2 转换为 ("Tom", 18), ("Same", 25) 这样的格式。


### Find操作
期望的调用方式是这样的：传入一个切片指针，查询的结果保存在切片中。
```go
s := geeorm.NewEngine("sqlite3", "gee.db").NewSession()
var users []User
s.Find(&users);
```

Find 的代码实现比较复杂，主要分为以下几步：

- 1)`destSlice.Type().Elem()` 获取切片的单个元素的类型 `destType`，使用 `reflect.New()` 方法创建一个 `destType` 的实例，作为 `Model()` 的入参，映射出表结构 `RefTable()`。
- 2）根据表结构，使用 `clause `构造出 `SELECT `语句，查询到所有符合条件的记录 `rows`。
- 3）遍历每一行记录，利用反射创建 `destType` 的实例 `dest`，将 `dest` 的所有字段平铺开，构造切片 `values`。
- 4）调用 `rows.Scan()` 将该行记录每一列的值依次赋值给 `values` 中的每一个字段。
- 5）将 `dest` 添加到切片 `destSlice` 中。循环直到所有的记录都添加到切片 `destSlice` 中。
