package main

func reverseWords(s string) string {
	runes := []rune(s)
	stack := []rune{}
	res := []rune{}
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == " " {
			for j := 0; j < len(stack); j++ {
				res = append(res, stack[len(stack) - j - 1])
			}
			res = append(res, runes[i])
			stack = nil
		} else {
			stack = append(stack, runes[i])
		}
	}

	for j := 0; j < len(stack); j++ {
		res = append(res, stack[len(stack) - j - 1])
	}

	return string(res)
}