package main

import (
	"fmt"
	"strings"

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

// User 用户
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

//拼接语句实现批量插入
func batchInsertUsers(users []*User) error {
	//存放 (?,?)的slice
	valueStrings := make([]string, 0, len(users))
	//存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	//遍历users准备相关数据
	for _, u := range users {
		//此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?,?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	//自行拼接要执行的具体语句
	stmt := fmt.Sprintf("insert into user (name, age) values %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

//BatchInsertUsers2 方式二
func BatchInsertUsers2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"insert into user (name,age) values (?),(?),(?)",
		users..., //如果args实现了 driver.Valuer,sqlx,In 会通过调用value()来展开
	)
	fmt.Println(query) //查看生成的querystring
	fmt.Println(args)  //查看生成的args
	_, err := db.Exec(query, args...)
	return err
}

//BatchInsertUsers3 方式3
func BatchInsertUsers3(users []*User) error {
	_, err := db.NamedExec("insert into user (name, age) values (:name,:age)", users)
	return err
}

func main() {
	initDB()
	u1 := User{Name: "托儿索", Age: 18}
	u2 := User{Name: "小学僧", Age: 21}
	u3 := User{Name: "儿童劫", Age: 19}

	// 方法1
	users := []*User{&u1, &u2, &u3}
	err := batchInsertUsers(users)
	if err != nil {
		fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	}

	// 方法2
	users2 := []interface{}{u1, u2, u3}
	err = BatchInsertUsers2(users2)
	if err != nil {
		fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	}

	// 方法3
	users3 := []*User{&u1, &u2, &u3}
	err3 := BatchInsertUsers3(users3)
	if err3 != nil {
		fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err3)
	}
}
