package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

var (
	message = "%d %s \n"
	option  uint8
)

func main() {
	for {
		menu()
		fmt.Printf("option value = %d\n", option)
		dirllController(option)
	}
}

func menu() {
	fmt.Println("go 交互练习--")
	fmt.Printf(message, 0, "退出")
	fmt.Printf(message, 1, "空格分割读取")
	fmt.Printf(message, 2, "指定结束读取")
	fmt.Print("Please enter your select:")
	_, err := fmt.Scanln(&option)
	errController(err)
}

func dirllController(option uint8) {
	switch option {
	case 0:
		os.Exit(1)
	case 1:
		stringRead()
	case 2:
		bufioDrill()
	default:
		errController(1)
	}
}

func stringRead() {
	fmt.Println("fmt.Sscan dirll.full string : my name is Qingwr . use message fmt.Sscan(full, &object, &n, &be, &target)")
	full := "my name is Qingwr"
	var object, n, be, target string
	fmt.Sscan(full, &object, &n, &be, &target)
	fmt.Printf("object=%s n=%s be=%s target=%s\n", object, n, be, target)
	fmt.Println("---------------------")
}

func bufioDrill() {
	fmt.Printf("bufio.NewReader(os.Stdin).ReadString('\\n') 指定结束符号读取\n")
	fmt.Print("Please enter any char:")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	errController(err)
	fmt.Printf("bufioDrill output :%s", input)
}

func errController(err interface{}) {
	if err != nil {
		fmt.Println("大哥,你又挖坑了!")
		fmt.Println("------------------")
		log.Fatal(err)
		menu()
	}
}
