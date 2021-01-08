package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func initDB() (err error) {
	//DNS:Data Source Name
	address := "root:123456@tcp(127.0.0.1:3306)/dbname"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", address)
	if err != nil {
		return err
	}
	//尝试与数据库建立连接(校验address是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
}
