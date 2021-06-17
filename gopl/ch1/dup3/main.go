package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//全部输入数据读到内存中，一次分割为多行，然后处理它们
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			//使用Fprintf与表示任意类型默认格式值的动词%v，向标准错误流打印一条信息
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}