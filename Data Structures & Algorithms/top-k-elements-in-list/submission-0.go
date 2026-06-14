func topKFrequent(nums []int, k int) []int {
	max := 0
    counts := make(map[int]int)
    for _, v := range nums {
        counts[v]++
        if counts[v] > max {
            max = counts[v]
        }
    }

    ans := make([][]int, max + 1)

    for key, v := range counts {
        ans[v] = append(ans[v], key)
    }

    count := 0
    ans2 := make([]int, 0)
    
    for i := len(ans) - 1; i > 0; i-- {
        if len(ans[i]) == 0 {
            continue
        }
        for _, value := range ans[i] {
            ans2 = append(ans2, value)
            count++
            if count == k {
                return ans2
            }
        }
    }
    
    return ans2
}
