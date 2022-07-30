package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	//连接驱动
	db, _ := sql.Open("sqlite3", "demo")
	defer func() { _ = db.Close() }()

	//创建表格
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")

	//插入数据
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}

	//查询数据
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}

	/**
	output:
	2022/07/30 17:11:33 2
	2022/07/30 17:11:33 Tom
	*/
}
