import "slices"
func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	global := make([][]int, 0)
	var dfs func (start int, res []int, count int)
	dfs = func (start int, res []int, count int) {
		if count == target {
			temp := make([]int, len(res))
			copy(temp, res)
			global = append(global, temp)
			return
		}

		if start >= len(candidates) || count > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i - 1] {
				continue
			}
			dfs(i+ 1, append(res, candidates[i]), count + candidates[i])
		}
	}
	dfs(0, []int{}, 0)
	return global
}
