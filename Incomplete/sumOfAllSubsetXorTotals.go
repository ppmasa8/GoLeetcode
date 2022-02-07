func subsetXORSum(nums []int) int {
    var xorSum int = 0

	for i := 0; i < (1 << len(nums)); i++ {
		var xor int = 0
		for j := 0; j < len(nums); j++ {
			if (i & (1 << j)) >= 1 {
				xor ^= nums[j]
			}
		}
		xorSum += xor
	}

	return xorSum
}
