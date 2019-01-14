/*
	输入流读取
	bufio.NewReader() 构造函数的签名为： func NewReader(rd io.Reader) *Reader。
*/
package main

import (
	"bufio"
	"os"
	"fmt"
)

var (
	inputReader *bufio.Reader //
	input       string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin) ////创建一个读取器，并将其与标准输入绑定。
	fmt.Print("Please enter some input: ")
	/*读取器对象提供一个方法 ReadString(delim byte) ，该方法从输入中读取内容
	 ，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。*/
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s", input)
 	}
}
