package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

// 连接数据库
func InitDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/Test_China_Region"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	fmt.Printf("连接数据库成功\n")
	return
}

type province struct {
	pid   string
	pname string
}
type city struct {
	cid   string
	pid   string
	cname string
}
type town struct {
	tid   string
	cid   string
	tname string
}

// 1. 查询广东省下所有的城市 下所有的乡镇.
func queryAlltownOfPronvince(p string) (res []string) {
	strSql := "SELECT town.tid from province , city , town where province.pid = city.pid and city.cid = town.cid and province.pid = ?;"
	rows, err := db.Query(strSql, p)

	if err != nil {
		fmt.Printf("Query all town of province failed , err :%s\n", err)
		return
	}
	var tmp string
	for rows.Next() {
		err = rows.Scan(&tmp)
		//fmt.Println("ttt ", tmp)
		if err != nil {
			fmt.Printf("Query a town of province failed , err :%s\n", err)
			return
		}
		res = append(res, tmp)
	}
	return res
}

//2. 查询某个乡镇属于哪个 省份
func queryTownOfProvince(t string) (res string) {
	strSql := "select p.pid from province p , city c , town t where p.pid = c.pid and c.cid = t.cid and t.tid = ? ;"

	err := db.QueryRow(strSql, t).Scan(&res)
	if err != nil {
		fmt.Printf("query failed , err : %s\n", err)
		return
	}
	fmt.Printf("Query Province success\n")
	return
}
func main() {

	err := InitDB()
	if err != nil {
		fmt.Printf("connect DB failed , err : %s\n", err)
		return
	}
	defer db.Close()

	//实现一个国家树 能够记录 中国 - 省份 - 城市 - 乡镇 的结果.  实现接口:
	//
	//1. 查询广东省下所有的城市 下所有的乡镇.

	//fmt.Println(queryAlltownOfPronvince("P002"))

	//2. 查询某个乡镇属于哪个 省份
	//fmt.Println(queryTownOfProvince("T2002_01"))

	fmt.Println(time.Now())
}
