package main

import (
	"net/http"
	"io"
	"os"
)

func main() {

	getDrill(os.Args[1])
}

func getDrill(addr string)  {
	resp, err := http.Get(addr)
	checkErr(err)
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}