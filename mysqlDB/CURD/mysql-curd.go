package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

type tbProjInfo struct {
	id   int
	name string
}

func initDB() (err error) {
	//DNS:Data Source Name
	address := "root:123456@tcp(127.0.0.1:3306)/test02"
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

//单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果
func queryRow() {
	sqlString := "select id,name from tb_proj_info where id = ?"
	var projInfo tbProjInfo
	err := db.QueryRow(sqlString, 1).Scan(&projInfo.id, &projInfo.name)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s\n", projInfo.id, projInfo.name)
}

//多行查询 db.Query()执行一次
func queryMultiRowDemo() {
	sqlStr := "select id,name from tb_proj_info where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	//循环读取结果集中的数据
	for rows.Next() {
		var projInfo tbProjInfo
		err := rows.Scan(&projInfo.id, &projInfo.name)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s\n", projInfo.id, projInfo.name)
	}
}

//插入数据
func insertRow() {
	sqlStr := "insert into tb_proj_info(name) values(?)"
	ret, err := db.Exec(sqlStr, "青草沙")
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, id=%d\n", theID)
}

//更新数据
func updateRow() {
	sqlStr := "update tb_proj_info set name = ? where id = ?"
	ret, err := db.Exec(sqlStr, "DAS", 2)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作受影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

//删除数据
func deleteRow() {
	sqlStr := "delete from tb_proj_info where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作受影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	queryRow()
	queryMultiRowDemo()
	insertRow()
	updateRow()
	deleteRow()
}
