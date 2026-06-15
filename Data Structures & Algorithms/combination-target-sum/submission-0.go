func combinationSum(nums []int, target int) [][]int {
    ans := make([][]int, 0)
	var dfs func (start int, curr []int, sum int)
	dfs = func (start int, curr []int, sum int) {
		if sum == target {
			temp := make([]int, len(curr))
            copy(temp, curr)
            ans = append(ans, temp)
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(nums); i++ {
			dfs(i, append(curr, nums[i]), sum + nums[i])
		}
	}
	dfs(0, []int{}, 0)
	return ans
}
