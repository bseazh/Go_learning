package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/goMySqltest"
	// 也可以使用MustConnect连接不成功就panic
	// Connect = Open + Ping 检验同时尝试去Ping
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	// 设置sql线程池最大容量
	db.SetMaxOpenConns(20)
	// 设置sql最大空余容量
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed , err %s\n", err)
		return
	}
	defer db.Close()
	fmt.Printf("连接数据库成功!\n")

	var u user
	sqlStr1 := "select id , name , age from user where id = 1 "
	db.Get(&u, sqlStr1)
	fmt.Printf("ID:%d,name:%s,age:%d\n", u.ID, u.Name, u.Age)

	var userlist []user
	sqlStr2 := "select id , name , age from user"
	db.Select(&userlist, sqlStr2)
	fmt.Printf("%#v\n", userlist)

}
