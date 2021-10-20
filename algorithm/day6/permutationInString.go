package main

func checkInclusion(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	v1 := make([]int, 26)
	v2 := make([]int, 26)
	for i := range s1 {
		v1[s1[i]-'a']++
		v2[s2[i]-'a']++
	}
	if match(v1, v2) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		v2[s2[i]-'a']++
		v2[s2[i-len(s1)]-'a']--
		if s2[i] != s2[i-len(s1)] && match(v1, v2) {
			return true
		}
	}
	return false
}

func match(s1, s2 []int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
