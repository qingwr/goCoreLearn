package main

import (
	"html/template"
	"fmt"
	"net/http"
	"io"
	"log"
)

type Person struct {
	Name string
	Age uint8
}

func main()  {

	http.HandleFunc("/", LogPanlcs(index)) //匹配不到路径采用
	http.HandleFunc("/if", LogPanlcs(htmlIfDrill))
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

func htmlIfDrill(w http.ResponseWriter, r *http.Request)  {
	p := Person{Name:"卿旺荣", Age:29}
	templateShow(w,"index.html", p)
}

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "hello go!")
}

func templateShow(w io.Writer, templateFile string, p interface{})  {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	if err := t.Execute(w, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}

// 错误处理
func LogPanlcs(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil{
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handlerFunc(writer, request)
	}
}