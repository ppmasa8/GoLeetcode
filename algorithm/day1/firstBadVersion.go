package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left + 1 < right {
		mid := (left+right)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid
		}
	}

	if isBadVersion(left) {
		return left
	}

	return right
}
