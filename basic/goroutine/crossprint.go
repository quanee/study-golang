package main

import "fmt"

var chanNum = make(chan int)
var chanAlp = make(chan string)


func printAlpha() {
	i := 0
	for {
		fmt.Println(<-chanAlp)
		chanNum <- i
		i++
	}
}

func printNum() {
	i := 0
	for {
		chanAlp <- string(rune('a' + i))
		i++
		fmt.Println(<-chanNum)
	}
}

func main() {
	go printAlpha()
	go printNum()
	for{;}
}
