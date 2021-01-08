package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test02?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed,err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

//NamedExec() 用来绑定SQL语句与结构体或map中的同名字段
func insertUser() (err error) {
	sqlStr := "insert into user (name,age) values(:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "jbf",
			"age":  24,
		})
	return
}

type tbProjInfo struct {
	ID   int
	Name string
}

func namedQuery() {
	sqlStr := "select * from tb_proj_info where `name`=:name"
	//使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "青草沙"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u tbProjInfo
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := tbProjInfo{
		Name: "贵州地灾",
	}
	// 使用结构体命名查询，根据结构体字段的db tag 进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NameQuery failed, err:%b\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u tbProjInfo
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func main() {
	initDB()
	// insertUser()
	namedQuery()
}
