package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	vals := req.URL.Query()
	sleep_time, _ := strconv.ParseInt(vals.Get("sleep"), 10, 64)
	time.Sleep(time.Duration(sleep_time) * time.Millisecond)
	fmt.Fprintln(w, "Hello World")
}

func main() {

	http.HandleFunc("/200", hello)

	http.ListenAndServe(":8090", nil)
}
