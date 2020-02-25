package main

import (
	"fmt"
	"os"
)

func main() {
	// 变量在声明时直接初始化，默认值是其类型的零值
	var s, sep  string
	// ':='：短变量声明，
	for i := 1; i < len(os.Args); i ++ {
		// '+' 运算符表示连接字符串
		s += sep + os.Args[i]
		sep	= " "
	}
	fmt.Println(s)
}
