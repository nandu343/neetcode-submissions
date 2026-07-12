func subsets(nums []int) [][]int {
	global := make([][]int,0)
	var dfs func (start int, res []int)
	dfs = func (start int, res []int) {
		if start == len(nums) {
			temp := make([]int, len(res))
			copy(temp, res)
			global = append(global, temp)
			return
		}
		dfs(start + 1, res)
		dfs(start + 1, append(res, nums[start]))
	}
	dfs(0, []int{})
	return global
}
