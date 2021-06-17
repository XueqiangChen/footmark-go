// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建空 map
	counts := make(map[string]int)
	// input type is bufio.Scanner
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//等价于：
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
		if input.Text() == "quit" {
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			// %d表示以十进制形式打印一个整型操作数，而%s则表示把字符串型操作数的值展开。
			//%d          十进制整数
			//%x, %o, %b  十六进制，八进制，二进制整数。
			//%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
			//%t          布尔：true或false
			//%c          字符（rune） (Unicode码点)
			//%s          字符串
			//%q          带双引号的字符串"abc"或带单引号的字符'c'
			//%v          变量的自然形式（natural format）
			//%T          变量的类型
			//%%          字面上的百分号标志（无操作数）
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
