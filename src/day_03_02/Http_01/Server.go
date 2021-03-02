package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./Http_showPage.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v\n", err)))
		return
	}
	w.Write(b)
}
func main() {
	http.HandleFunc("/Path1", f1)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
