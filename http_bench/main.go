package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	url := "http://localhost:8090/200?sleep=200"
	method := "GET"

	client := &http.Client{}
	req1, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(body))
}

func main() {

	http.HandleFunc("/testing", hello)

	http.ListenAndServe(":8080", nil)
}
