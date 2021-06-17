package main

import (
	"fmt"
	"os"
)

// 1.1 课后习题
func main() {
	s, sep := "", ""
	for idx, arg := range os.Args[0:] {
		if idx != 0 {
			fmt.Println(idx, ":", arg)
		}
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
