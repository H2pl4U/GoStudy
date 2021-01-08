package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type tbProjInfo struct {
	ID   int
	Name string
}

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

//查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id,name from tb_proj_info where id = ?"
	var info tbProjInfo
	err := db.Get(&info, sqlStr, 3)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s \n", info.ID, info.Name)
}

//查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id,name from tb_proj_info where id > ?"
	var infos []tbProjInfo
	err := db.Select(&infos, sqlStr, 0)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("projInfo:%v \n", infos)
}

//插入数据
func insertRow() {
	sqlStr := "insert into tb_proj_info (name) values (?)"
	ret, err := db.Exec(sqlStr, "柳州水利")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, theID is %d. \n", theID)
}

//更新数据
func updateRow() {
	sqlStr := "update tb_proj_info set name = ? where id = ?"
	ret, err := db.Exec(sqlStr, "小寺沟", 4)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //受影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, theID is %d. \n", n)
}

//删除数据
func deleteRow() {
	sqlStr := "delete from tb_proj_info where id = ?"
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //受影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, theID is %d. \n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	queryRowDemo()
	queryMultiRowDemo()
	// insertRow()
	// updateRow()
	// deleteRow()
}
