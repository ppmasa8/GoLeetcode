func lemonadeChange(bills []int) bool {
    fifths, tens := 0, 0

	for _, v := range bills {
		if v == 5 {
			fifths++
		} else if v == 10 {
			fifths--
			tens++
		} else if tens > 0 {
			tens--
			fifths--
		} else {
			fifths -= 3
		}
		if fifths < 0 {
			return false
		}
	}
	return true
}
