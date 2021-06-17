// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Println(resp.Status)
	}
}