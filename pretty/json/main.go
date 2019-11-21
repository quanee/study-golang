package format

var SPACE = "    "

func Format(rawjson string) string {
	var result []rune

	json := []rune(rawjson)
	length := len(json)
	number := 0

	for i := 0; i < length; i++ {
		key := json[i]
		if key == '[' || key == '{' {
			if i-1 > 0 && json[i-1] == ':' {
				result = append(result, '\n')
				for _, x := range indent(number) {
					result = append(result, x)
				}
			}
			result = append(result, rune(key))
			result = append(result, '\n')

			number++
			for _, x := range indent(number) {
				result = append(result, x)
			}

			continue
		}

		if key == ']' || key == '}' {
			result = append(result, '\n')

			number--
			for _, x := range indent(number) {
				result = append(result, x)
			}
			result = append(result, rune(key))

			if (i+1) < length && json[i+1] != ',' {
				result = append(result, '\n')
			}
			continue
		}

		if key == ',' {
			result = append(result, rune(key))
			result = append(result, '\n')
			for _, x := range indent(number) {
				result = append(result, x)
			}
			continue
		}
		result = append(result, rune(key))
	}

	return string(result)
}

func indent(number int) []rune {
	var result []rune
	for i := 0; i < number; i++ {
		for _, x := range []rune(SPACE) {
			result = append(result, x)
		}
	}
	return result
}
