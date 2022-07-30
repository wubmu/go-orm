package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"orm"
)

func main() {
	engine, _ := orm.NewEngine("sqlite3", "demo.db")
	defer engine.Close()

	//s := engine.NewSession()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)

	/*
		[info ] 2022/07/30 18:09:22 orm.go:26: Connect database success
		[info ] 2022/07/30 18:09:22 raw.go:42: DROP TABLE IF EXISTS User;  []
		[info ] 2022/07/30 18:09:22 raw.go:42: CREATE TABLE User(Name text);  []
		[info ] 2022/07/30 18:09:22 raw.go:42: CREATE TABLE User(Name text);  []
		[error] 2022/07/30 18:09:22 raw.go:44: table User already exists
		[info ] 2022/07/30 18:09:22 raw.go:42: INSERT INTO User(`Name`) values (?), (?)  [Tom Sam]
		Exec success, 2 affected
		[info ] 2022/07/30 18:09:22 orm.go:34: Close database success
	*/
}
