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

func main() {
	// dsn : Data Source Name
	// 格式 : "用户名:密码@tcp(本机IP地址,以'.'分割)/数据库名称"
	dsn := "root:root@tcp(192.168.43.160:3306)/goMySqltest"

	// 只校验dsn的格式是否正确,不会校验其用户名和密码
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn %s invaild , err : %v\n", dsn, err)
		return
	}
	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Printf("Open failed , err : %v\n", err)
		return
	}
	// 打开之后一定记得关闭数据库
	defer db.Close()
	fmt.Println("数据库连接成功")
}
