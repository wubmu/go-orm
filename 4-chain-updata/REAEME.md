## 链式操作与更新删除

- 通过链式(chain)操作，支持查询条件(where, order by, limit 等)的叠加。
- 实现记录的更新(update)、删除(delete)和统计(count)功能。

### 链式调用
某个对象调用某个方法后，将该对象的引用/指针返回，即可以继续调用该对象的其他方法。通常来说，当某个对象需要一次调用多个方法来设置其属性时，就非常适合改造为链式调用了。

SQL 语句的构造过程就非常符合这个条件。SQL 语句由多个子句构成，典型的例如 SELECT 语句，往往需要设置查询条件（WHERE）、限制返回行数（LIMIT）等。理想的调用方式应该是这样的：
```go
s := geeorm.NewEngine("sqlite3", "gee.db").NewSession()
var users []User
s.Where("Age > 18").Limit(3).Find(&users)
```

，`WHERE`、`LIMIT`、`ORDER BY` 等查询条件语句非常适合链式调用。

### 3 First只返回一条数据
我们期望 SQL 语句只返回一条记录，比如根据某个童鞋的学号查询他的信息，返回结果有且只有一条。结合链式调用，我们可以非常容易地实现 First 方法。