package main

func sortedSquares(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		res = append(res, v*v)
	}
	return quickSort(res)
}

func quickSort(nums []int) (res []int) {
	if len(nums) < 2 {
		return nums
	}

	// shift
	pivot := nums[0]
	nums = nums[1:]

	big, small := []int{}, []int{}

	for _, v := range nums {
		if v < pivot {
			small = append(small, v)
		} else {
			big = append(big, v)
		}
	}

	small = quickSort(small)
	big = quickSort(big)

	res = append(small, pivot)
	res = append(res, big...)

	return
}
