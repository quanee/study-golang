package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x)-1; i++ {
		tmp := rand.Intn(len(x)-i)
		x[len(x)-i-1], x[tmp] = x[tmp], x[len(x)-i-1]
	}
	fmt.Println(x)
	wg := sync.WaitGroup{}
	wg.Done()
}
