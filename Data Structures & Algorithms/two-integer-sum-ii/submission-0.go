func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1
	for l < r {
		if numbers[l] + numbers[r] == target {
			return []int{l + 1, r + 1}
		}

		if numbers[l] + numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return nil
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
