func rob(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}
	
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		}
		return nums[0]
	}
	max := make([]int, len(nums) +1 )
	max[0] = 0
	max[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		max[i] = comp(max[i - 1], max[i - 2] + nums[i - 1])
	}

	return max[len(nums)]

}


func comp (a int, b int) int {
	if a > b {
		return a
	}
	return b
}
