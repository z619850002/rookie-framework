package db

import (
	"testing"
)

var ds *DataSource

func initMysql() {
	mysqlconfig := &MysqlConfig{
		MysqlAddr:     "localhost",
		MysqlPort:     "3306",
		MysqlUser:     "root",
		MysqlPassword: "12345678",
		MysqlDB:       "test111",
	}
}

type User struct {
	Name string
	Age  int
}

func TestExec(t *testing.T) {
	initMysql()
	res := []User{}
	rows, _ := ds.Query("select name,age from table1")
	for rows.Next() {
		r := User{}
		rows.Scan(&r.Name, &r.Age)
		res = append(res, r)
	}
	if res[0] != (User{Name: "lihao", Age: 25}) {
		t.Error("Exec函数测试错误！")
	}
}

func TestQuery(t *testing.T) {
	initMysql()
	//for UPDATE operation
	ds.Exec(`update table1 set age=? where name=?`, 100, "lihao") //update to check
	res := []User{}
	rows, _ := ds.Query("select name,age from table1")
	for rows.Next() {
		r := User{}
		rows.Scan(&r.Name, &r.Age)
		res = append(res, r)
	}
	if res[0] != (User{Name: "lihao", Age: 100}) {
		t.Error("Query函数测试错误！")
	}
	ds.Exec(`update table1 set age=25 where name="lihao"`) //restore the data

	//for INSERT operation
	ds.Exec(`insert into table1 (name,age) values ("tianjun",21)`)
	res = []User{}
	rows, _ = ds.Query(`select name,age from table1 where name="tianjun"`)
	for rows.Next() {
		r := User{}
		rows.Scan(&r.Name, &r.Age)
		res = append(res, r)
	}
	if res[0] != (User{Name: "tianjun", Age: 21}) {
		t.Error("Query函数测试错误！")
	}

	ds.Exec(`delete from table1 where name="tianjun"`) //delete tianjun
	//for DELETE operation, check wether tianjun is deleted
	res = []User{}
	rows, _ = ds.Query(`select name,age from table1 where name="tianjun"`)
	for rows.Next() {
		r := User{}
		rows.Scan(&r.Name, &r.Age)
		res = append(res, r)
	}
	if len(res) != 0 {
		t.Error("Query函数测试错误！")
	}
}
