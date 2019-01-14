/*
	命令行参数读取
	use message: go run args.go 1 2 3 4
	os.Args[0]为自身
*/
package main

import (
	"os"
	"fmt"
)

func main()  {

	for i,val := range os.Args{
		fmt.Printf("index:%d, value:%s\n", i, val)
	}
}
