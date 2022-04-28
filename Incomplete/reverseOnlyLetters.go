func reverseOnlyLetters(S string) string {
	ss := []byte(S)

	for i, j := 0, len(ss)-1; i<j; {
		if isAlphabet1(ss[i]) && isAlphabet1(ss[j]) {
			ss[i], ss[j] = ss[j], ss[i]
			i++
			j--
			continue
		}
		if !isAlphabet1(ss[i]) && !isAlphabet1(ss[j]) {
			i++
			j--
			continue
		}
		if isAlphabet1(ss[i]) {
			j--
			continue
		}
		i++
	}

	return string(ss)
}

func isAlphabet1(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
