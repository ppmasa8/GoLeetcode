package main

func moveZeroes(nums []int)  {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cnt] = nums[i]
			cnt++
		}
	}

	for j := cnt; j < len(nums); j++ {
		nums[j] = 0
	}
}
