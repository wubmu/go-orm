## day1-database_sql
项目目录
```
day1-database-sql/
    |--log/          # 日志
        |--log.go
    |--session/      # 数据库交互
        |--raw.go
    |--orm.go     # 用户交互
    |--go.mod
```

	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
session中
raw
- Session: 结构体包含三个成员变量，
  - `db *sql.DB`
  - `sql strings.Builder`
  - `sqlVal []interface{}`
- Session的方法：封装的目的：1.统一打印日志，2二是执行完成后，
清空 (s *Session).sql 和 (s *Session).sqlVars 两个变量。这样 
Session 可以复用，开启一次会话，可以执行多次 SQL。
  - Exec()
  - Query()
  - QueryRow()