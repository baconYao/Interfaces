package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
)

// 自己寫一個Writer interface
type logWriter struct {}

func getGooglePage() {
	// https://golang.org/pkg/net/http/#Client.Get
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// 第一種印出 body 的方式
	// make: 用來宣告一個有範圍限制的 slice
	// 為何這裡要用 []byte? 因為 resp 是 *Response type  https://golang.org/pkg/net/http/#Response
	// 從 resp 文件裡面找到 Body 會看到 Body io.ReadCloser，表示 Body 的 type 是 io.ReadCloser interface https://golang.org/pkg/io/#ReadCloser
	// 從 ReadCloser interface 有 Reader 和 Closer 兩個 interface
	// 最後在 Reader 這個 interface 看到 它定義的 method -> Read(p []byte) (n int, err error)
	// 可以發現 Read 需要 []byte type 的參數。因此我們需要給予 []byte 到 resp.Body.Read()
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@")

	// 第二種印出 body 的方式
	// https://golang.org/pkg/io/#Copy
	// https://golang.org/pkg/io/#Writer
	// https://golang.org/pkg/os/#File
	// io.Copy(os.Stdout, resp.Body)

	// 第三種印出 body 的方式，自己寫 Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}