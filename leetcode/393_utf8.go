package main

func main() {

}

func validUtf8(data []int) bool {
	flag := 0
	for _, d := range data {
		if flag == 0 {
			if d>>7 == 0 {
				continue
			} else if d>>5 == 6 {
				flag = 1
			} else if d>>4 == 14 {
				flag = 2
			} else if d>>3 == 30 {
				flag = 3
			}
		} else if flag > 0 {
			if d>>6 == 2 {
				flag--
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
