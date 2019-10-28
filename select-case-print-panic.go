package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
	// 两种输出结果 select 的case维护一个堆从所有符合条件中随机选取一个
	/*
	1
	*/
	/*
	panic: hello

	goroutine 1 [running]:
	main.main()
		E:/goyard/src/study-golang/main.go:18 +0x20f
	*/
}
