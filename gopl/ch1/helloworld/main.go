package main

// import声明必须跟在文件的package声明之后。
// 随后，则是组成程序的函数、变量、常量、类型的声明语句（分别由关键字func、var、const、type定义）
import "fmt"

// main包比较特殊。它定义了一个独立可执行的程序，而不是一个库。
// 在main里的main 函数 也很特殊，它是整个程序执行时的入口
func main(){
	fmt.Println("Hello, 世界")
}
