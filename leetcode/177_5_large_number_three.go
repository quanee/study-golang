package main

import "fmt"

func main() {
	digits := []int{8, 6, 7, 1, 0}
	fmt.Println(largestMultipleOfThree(digits))
}

func largestMultipleOfThree(digits []int) string {
	sum := 0
	for _, v := range digits {
		sum += v
	}
	if sum == 0 {
		return "0"
	}
	if sum%3 != 0 {
		return ""
	}
	res := ""
	quick_sort(digits, 0, len(digits)-1)
	for i := len(digits) - 1; i >= 0; i++ {
		res += string(digits[i])
	}

	return res
}

func swap(a int, b int) (int, int) {
	return b, a
}

func partition(aris []int, begin int, end int) int {
	pvalue := aris[begin]
	i := begin
	j := begin + 1
	for j < end {
		if aris[j] < pvalue {
			i++
			aris[i], aris[j] = swap(aris[i], aris[j])
		}
		j++
	}
	aris[i], aris[begin] = swap(aris[i], aris[begin])
	return i
}

func quick_sort(aris []int, begin int, end int) {
	if begin+1 < end {
		mid := partition(aris, begin, end)
		quick_sort(aris, begin, mid)
		quick_sort(aris, mid+1, end)
	}
}
