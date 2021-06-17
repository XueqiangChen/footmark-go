package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// 程序的命令行参数可从os包的Args变量获取；
	// os包外部使用os.Args访问该变量。
	// os.Args变量是一个字符串（string）的切片（slice）
	// os.Args的第一个元素：os.Args[0]，是命令本身的名字；其它的元素则是程序启动时传给它的参数。
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
