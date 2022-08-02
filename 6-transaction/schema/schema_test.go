package schema

import (
	"orm/dialect"
	"testing"
)

type User struct {
	Name string `orm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}

func TestSchema_RecordValues(t *testing.T) {
	//u1 := &User{Name: "Tom", Age: 18}
	//schema := Parse(u1, TestDial)

}
