//https://studygolang.com/articles/14669?fr=sidebar
package main

import (
	"flag"
	"os"
	"log"
	"bufio"
	"fmt"
)

func main()  {
	//blockRead()
	rowRead()
}

/*
	分段读取
	!!!注意读取长度,分隔符也算字符数,容易出现乱码
*/
func blockRead()  {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println(*fptr)

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

/*
	逐行读取
*/
func rowRead()  {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("read file:",*fptr)

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil{
		log.Fatal(err)
	}
}
