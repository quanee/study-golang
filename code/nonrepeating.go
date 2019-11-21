package main

import "fmt"

var lastOccured = make([]int, 0xffff)
func lengthOfNonRepeatingSubStr(s string) int {
	// lastOccured := make(map[rune]int)
	for i := range lastOccured{
		lastOccured[i] = 0
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccured[ch]; lastI > start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i-start+1
		}
		lastOccured[ch] = i+1
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
}
