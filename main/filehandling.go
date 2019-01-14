//https://studygolang.com/articles/10552
package main

import (
	"io/ioutil"
	"fmt"
	"flag"
	"github.com/gobuffalo/packr"
)

func main() {
	//#ioutilRead(os.Args[1])
	//coommandFlag()
	//packrRead()
	ioutilWrite()
}


/*
	ioutil读取
*/
func ioutilRead(file string)  {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("reading %s err:%v", file, err)
	}
	fmt.Println("Contents of file:", string(data))
}

/*
	ioutil写
*/
func ioutilWrite(){
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("test.txt", d1, 0644)
	check(err)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}


/*
获取文件路径1
命令参数绑定
	usage
> go run filehandling.go -fpath=/path-of-file/test.txt
	value of fpath is  /path-of-file/test.txt
*/
func coommandFlag()  {
	fmt.Println("")
	fptr := flag.String("fpath", "test.txt", "file path to read from ")
	flag.Parse()
	fmt.Println("value of fpath is ", *fptr)
}

/*  获取文件路径2
	将文件绑定在二进制文件中
	go get -u github.com/gobuffalo/packr/...
*/
func packrRead()  {
	box := packr.NewBox("../main")
	data := box.String("test.txt")
	fmt.Printf("contents of file:\n%s", data)
}

