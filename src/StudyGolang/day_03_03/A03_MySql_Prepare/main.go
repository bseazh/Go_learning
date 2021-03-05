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

// MySql 多行查询
func queryMulitRow(id_Lowest int) {
	sqlStr := "select id,name,age from user where id >= ? "

	// Query 返回多行Rows
	rows, err := db.Query(sqlStr, id_Lowest)
	if err != nil {
		fmt.Printf("QueryMulitRow failed err : %v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("Rows Scan failed , err : %v\n", err)
			return
		}
		fmt.Printf("Id : %d , name : %s , age : %d\n", u1.id, u1.name, u1.age)
	}
}

// MySQL 预处理演示(插入语句)
func prepareInsert() {
	// 编写 需预处理批量的处理的语句
	sqlStr := "insert into user(name,age) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare Insert Rows failed , err : %s\n", err)
		return
	}
	// 记住一定要释放
	defer stmt.Close()
	// 自定义数据进行插入
	var m = map[string]int{
		"Charles": 21,
		"Mark":    34,
		"Vincent": 56,
		"Jseph":   62,
	}

	for key, value := range m {
		stmt.Exec(key, value)
	}
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
	prepareInsert()
	queryMulitRow(0)
}
