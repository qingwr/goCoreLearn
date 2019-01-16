package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main()  {
	resp, err := http.Get("https://www.baidu.com/")
	checkErr(err)
	io.Copy(os.Stdout,resp.Body)
	defer resp.Body.Close()

}

func checkErr(err error)  {
	if err != nil {
		fmt.Println("get err:", err)
		os.Exit(-1)
	}
}