package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {

	http.HandleFunc("/", LogPanlcs(hello)) //匹配不到路径采用
	http.HandleFunc("/user/login", LogPanlcs(login))
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("handle hello")
	fmt.Fprintln(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintln(w, "login")
}
