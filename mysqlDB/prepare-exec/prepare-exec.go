package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type tbProjInfo struct {
	id   int
	name string
}

func initDB() (err error) {
	//DNS:Data Source Name
	address := "root:123456@tcp(127.0.0.1:3306)/test02"
	// 不要使用:=，给全局变量赋值，然后在main函数中使用全局变量db
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

func prepareQueryDemo(id int) {
	sqlStr := "select id,name from tb_proj_info where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	//循环读取结果集中的数据
	for rows.Next() {
		var info tbProjInfo
		err := rows.Scan(&info.id, info.name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s\n", info.id, info.name)
	}
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	prepareQueryDemo(1)
}
