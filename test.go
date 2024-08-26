package main

func test(s string) bool {
	param := []rune{}
	pairs := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, char := range s {
		if len(param) == 0 {
			_, ok := pairs[char]

			if !ok {
				return false
			}
			param = append(param, char)
		} else {
			x := param[len(param)-1]
			val, ok := pairs[char]

			if ok {
				param = append(param, char)
			} else {
				val, ok = pairs[x]
				if val != char {
					return false
				}
				param = param[:len(param)-1]
			}
		}
	}
	return len(param) == 0
}
