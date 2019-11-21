package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	userCount := math.MaxInt64
	for i := 0; i < userCount; i++ {
		go func(i int) {
			// 做一些各种各样的业务逻辑处理
			fmt.Printf("go func: %d\n", i)
			time.Sleep(time.Second)
		}(i)
	}
}
