## 6 事务机制
- 介绍数据库中的事务(transaction)。
- 封装事务，用户自定义回调函数实现原子操作

### 事务的ACID
> 数据库事务(transaction)是访问并可能操作各种数据项的一个数据库操作序列，
> 这些操作要么全部执行,要么全部不执行，是一个不可分割的工作单位。事务由事务开始与事务结束之间执行的全部数据库操作组成。


如果一个数据库支持事务，那么必须具备 ACID 四个属性。

- 1）原子性(Atomicity)：事务中的全部操作在数据库中是不可分割的，要么全部完成，要么全部不执行。
- 2）一致性(Consistency): 几个并行执行的事务，其执行结果必须与按某一顺序 串行执行的结果相一致。
- 3）隔离性(Isolation)：事务的执行不受其他事务的干扰，事务执行的中间结果对其他事务必须是透明的。
- 4）持久性(Durability)：对于任意已提交事务，系统必须保证该事务对数据库的改变不被丢失，即使数据库出现故障。

### SQLite和go中的事务
```sql
sqlite> BEGIN;
sqlite> DELETE FROM User WHERE Age > 25;
sqlite> INSERT INTO User VALUES ("Tom", 25), ("Jack", 18);
sqlite> COMMIT;
```

Go 语言标准库 database/sql 提供了支持事务的接口。用一个简单的例子，看一看 Go 语言标准是如何支持事务的。

```go
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "demo.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS User(`Name` text);")

	tx, _ := db.Begin()
	_, err1 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Tom")
	_, err2 := tx.Exec("INSERT INTO User(`Name`) VALUES (?)", "Jack")
	if err1 != nil || err2 != nil {
		_ = tx.Rollback()
		log.Println("Rollback", err1, err2)
	} else {
		_ = tx.Commit()
		log.Println("Commit")
	}
}
```

Go 语言中实现事务和 SQL 原生语句其实是非常接近的。
调用 `db.Begin()` 得到 `*sql.Tx` 对象，使用 `tx.Exec()` 执行一系列操作，
如果发生错误，通过 `tx.Rollback()` 回滚，如果没有发生错误，则通过 `tx.Commit()` 提交。

### ORM支持事务
GeeORM 之前的操作均是执行完即自动提交的，每个操作是相互独立的。
之前直接使用 `sql.DB` 对象执行 SQL 语句，如果要支持事务，需要更改
为 `sql.Tx` 执行。在 `Session` 结构体中新增成员变量 `tx *sql.Tx`，
当 `tx` 不为空时，则使用 `tx` 执行 SQL 语句，否则使用 db 执行 SQL 语句。
这样既兼容了原有的执行方式，又提供了对事务的支持。