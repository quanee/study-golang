package main

func main() {

}

func findRepeatNumber(nums []int) int {
	hash := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := hash[v]; ok {
			return v
		}
		hash[v] = struct{}{}
	}
	return 0
}

func findRepeatNumber1(nums []int) int {
	hash := make([]int, len(nums))

	for _, v := range nums {
		if hash[v] == 1 {
			return v
		}
		hash[v] = 1
	}
	return 0
}
