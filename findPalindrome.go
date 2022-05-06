package main

import "fmt"

func main() {
	var str = "ffkffsshsubshshskfidddcabac"
	fmt.Println(findAllPalindrome(str))
}

func findPalindrome(str string, left int, right int) []string {
	result := []string{}
	for 0 <= left && right <= len(str) - 1 {
		if str[left] != str[right] {
			break
		}
		result = append(result, str[left:right+1])
		left -= 1
		right += 1
	}
	return result
}

func findAllPalindrome(str string) []string {
	result := []string{}
	lenStr := len(str)
	if lenStr == 0 {
		return result
	}
	if lenStr == 1 {
		result = append(result, str)
	}

	for i := 0; i < lenStr-1; i++ {
		s := findPalindrome(str, i-1, i+1)
		for i := 0; i < len(s); i++ {
			result = append(result, s[i])
		}
		s = findPalindrome(str, i-1, i)
		for i := 0; i < len(s); i++ {
			result = append(result, s[i])
		}
	}

	return result
}