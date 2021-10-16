package main

func rotate(nums []int, k int) {
	left := nums[len(nums) - k%len(nums):len(nums)]
	right := nums[0: len(nums) - k%len(nums)]
	copy(nums, append(left, right...))
}
