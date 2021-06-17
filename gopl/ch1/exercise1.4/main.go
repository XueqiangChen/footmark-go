// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
练习 1.4： 修改dup2，出现重复的行时打印文件名称。
 */
func main() {
	//map是一个由make函数创建的数据结构的引用
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），
		//被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fileName, content := getNameAndContent(line)
			fmt.Printf("%s: %d\t%s\n", fileName, n, content)
		}
	}
}

func getNameAndContent(context string) (fileName string, content string) {
	strs := strings.Split(context, "|")
	return strs[0], strs[1]
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()+"|"+input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}