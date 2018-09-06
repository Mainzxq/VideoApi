package main

import (
	"fmt"

	"./mock"
	"./real"
)

// 接口的定义和实现
// 由使用者定义接口

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {

	var r Retriever
	r = mock.Retriever{"this is fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
