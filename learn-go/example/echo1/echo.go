package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 变量在声明时直接初始化，默认值是其类型的零值
	//var s, sep  string
	// ':='：短变量声明，
	//for i := 1; i < len(os.Args); i ++ {
	//	// '+' 运算符表示连接字符串
	//	s += sep + os.Args[i]
	//	sep	= " "
	//}
	//for _, arg := range(os.Args[1:]){
	//	s += sep + arg
	//	sep = " "
	//}
	fmt.Println(strings.Join(os.Args[1:], " "))
}
