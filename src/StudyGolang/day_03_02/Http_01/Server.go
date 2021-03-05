package main

// http server
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// Http_showPage.html 是在当前工作目录下编写好的html
	b, err := ioutil.ReadFile("./Http_showPage.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v\n", err)))
		return
	}
	w.Write(b)
}
func main() {
	//访问其目录时返回对应的函数实现 "响应"
	http.HandleFunc("/Path1", f1)
	//监听并服务 对应的IP地址
	http.ListenAndServe("127.0.0.1:9000", nil)
}
