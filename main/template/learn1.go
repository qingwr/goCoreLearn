package main

import (
	"html/template"
	"os"
	"fmt"
)

type Person struct{
	Name string
	Age int
}

func main() {
	tmpl, err := template.New("name").Parse("{{.Name}}")
	if err != nil {
		fmt.Printf("template install failed ,err:%v\n", err)
		return
	}
	tmpl.Execute(os.Stdout, Person{Name:"卿旺荣"})
	fmt.Println(tmpl.Name())
}