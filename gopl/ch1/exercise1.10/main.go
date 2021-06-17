// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递
		go fetchTwoTimes(url, ch)
		//go fetch(url, ch) // start a goroutine
	}
	// 这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出
	for range os.Args[1:] {
		// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
		//\直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchTwoTimes(url string, ch chan string) {
	ch <- "url:" + url + "\nresp1: " + fetch(url) + "\n" + "resp2: " + fetch(url)
}

func fetch(url string) string{
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprint(err) // send to channel ch
	}
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中（译注：可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据），
	// 因为我们需要这个方法返回的字节数，但是又不想要其内容。
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		return fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	return fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}