package main

// MySql 预处理

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// init()
)

var db *sql.DB // 连接池的一个对象
// 定义数据库中的元素
type user struct {
	id   int
	age  int
	name string
}

// 连接数据库 函数形式
func initDB() (err error) {
	// dsn : Data Source Name
	// 格式 : "用户名:密码@tcp(本机IP地址,以'.'分割)/数据库名称"
	dsn := "root:root@tcp(192.168.43.160:3306)/goMySqltest"
	// 只校验dsn的格式是否正确,不会校验其用户名和密码
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// MySQL 事务操作
func transaction() {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("DB Begin failed , err : %s\n", err)
		return
	}

	// 2. 编写 sqlStr1 , sqlStr2 语句
	sqlStr1 := "update user set age = age - 2 where id = 1 "
	sqlStr2 := "update user set age = age + 2 where id = 2 "

	// 逐一运行
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		fmt.Printf("sqlStr1 error : %s\n", err)
		tx.Rollback()
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Printf("sqlStr2 error : %s\n", err)
		tx.Rollback()
	}
	// 3. 提交事务
	err = tx.Commit()
	if err != nil {

		fmt.Printf("Commit failed , err %s\n", err)
		tx.Rollback()
	}
	// 4. 成功提交打印提示词
	fmt.Printf("Transaction Commit Success!\n")
}

// MySql 单行查询
func queryRow(id int) {
	sqlStr := "select id,name,age from user where id = ? "

	var u1 user
	// QueryRow 返回的 Row 是待执行的 sql语句
	// Scan 一旦执行完 即可 释放其资源
	err := db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	if err != nil {
		fmt.Printf("QueryOnce failed , err : %v\n", err)
		return
	}
	fmt.Printf("Id : %d , name : %s , age : %d\n", u1.id, u1.name, u1.age)
}
func main() {

	// 连接数据库
	err := initDB()
	if err != nil {
		fmt.Printf("Connect DB failed , err : %v\n", err)
		return
	}
	fmt.Println("数据库连接成功")
	//一定要记住关闭
	defer db.Close()
	transaction()
	queryRow(1) // 用于检验
	queryRow(2) // 用于检验
}
