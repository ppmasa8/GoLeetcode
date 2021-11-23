func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	mapping := map[byte]byte{}
	mapped := map[byte]interface{}{}
	for i := 0; i < len(s); i++ {
		if _, exist := mapping[s[i]]; !exist {
			if _, exist := mapped[t[i]]; exist {
				return false
			}
			mapping[s[i]] = t[i]
			mapped[t[i]] = nil
		} else {
			if mapping[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}
