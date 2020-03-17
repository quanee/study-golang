package main

func main() {

}

func countCharacters(words []string, chars string) int {
	count := [26]int{}
	for i := 0; i < len(chars); i++ {
		count[chars[i]-'a']++
	}
	res := 0
	flag := false
	temp := 0
	for i := 0; i < len(words); i++ {
		flag = false
		temp = len(words[i]) - 1
		for j := 0; j < len(words[i]); j++ {
			count[words[i][j]-'a']--
			if count[words[i][j]-'a'] < 0 {
				flag = true
				temp = j
				break
			}
		}
		for j := 0; j <= temp; j++ {
			count[words[i][j]-'a']++
		}
		if flag == false {
			res += len(words[i])
		}
	}

	return res
}
