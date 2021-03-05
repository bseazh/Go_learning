package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f2(w http.ResponseWriter, r *http.Request) {

	// 对于Get请求，参数都放在url上(Query param),请求体是没有数据的
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	age := queryParams.Get("age")

	fmt.Printf("name : %v , age : %v\n", name, age)

	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) // 我在服务端打印客户端发来的请求的Body

	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/Path2/", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)

}
