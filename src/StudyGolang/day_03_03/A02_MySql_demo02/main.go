package main

// MySql 建立连接
// Mac版本终端中打开：/usr/local/MySQL/bin/mysql -u root -p
// 192.168.43.160
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

// MySql 插入操作

func insertRow(userName string, userAge int) {
	sqlStr := "insert into user(name,age) values (?,?)"
	res, err := db.Exec(sqlStr, userName, userAge)
	if err != nil {
		fmt.Printf("Insert failed , err : %s\n", err)
		return
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get Lastid failed , err : %s\n", err)
		return
	}
	fmt.Printf("Insert Success , The LastInsertID : %v\n", lastId)
	// 检验结果 , queryRow(int(lastId))
}

// MySql 更新操作
func updateRow(userID int, userAge int) {
	sqlStr := "update user set age = ? where id = ? "
	res, err := db.Exec(sqlStr, userAge, userID)
	if err != nil {
		fmt.Printf("update failed err : %s\n", err)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row affected failed , err : %s\n", err)
		return
	}
	fmt.Printf("Update Success , Affected Row %d\n", rows)
	//queryRow(userID) //检验结果
}

// MySql 删除操作
func deleteRow(userID int) {
	sqlStr := "delete from user where id = ? "
	res, err := db.Exec(sqlStr, userID)
	if err != nil {
		fmt.Printf("delete Row failed , err : %s\n", err)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Row affected failed , err : %s\n", err)
		return
	}
	fmt.Printf("Delete Row Success , Affect Rows %d\n", rows)
	queryMulitRow(0) //检验结果
}
func main() {

	// 连接数据库
	err := initDB()
	if err != nil {
		fmt.Printf("Connect DB failed , err : %v\n", err)
		return
	}
	fmt.Println("数据库连接成功")

	// 单行查询
	//queryRow(2)

	// 多行查询
	//queryMulitRow(1)

	// 插入操作
	//insertRow("William", 53)

	// 更新操作
	//updateRow(7, 30)

	// 删除操作
	//deleteRow(7)
}
