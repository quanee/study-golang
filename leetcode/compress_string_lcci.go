package main

import (
	"fmt"
	"strconv"
)

func main() {
	//x := rune('1') - 1
	//xs := rune(x + 3)
	//fmt.Printf("%v%[1]q", xs)
	fmt.Println(compressString("abbccd"))
	//strconv.Itoa()
}

func compressString(S string) string {
	if len(S) < 3 {
		return S
	}
	res := make([]byte, 0)
	cur := S[0]
	count := 0
	res = append(res, S[0])
	for i := 0; i < len(S); {
		count = 0
		for i < len(S) && S[i] == cur {
			count++
			i++
		}
		res = append(res, []byte(strconv.Itoa(count))...)
		if i >= len(S) {
			break
		}
		res = append(res, S[i])
		cur = S[i]
	}
	if len(res) >= len(S) {
		return S
	}

	return string(res)
}
