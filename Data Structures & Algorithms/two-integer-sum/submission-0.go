func twoSum(nums []int, target int) []int {
    set := make(map[int]int)
    for i, v := range nums {
        k := target - v
        if val, e := set[k]; e == true {
            return []int{val, i}
        }
        set[v] = i
    }
    return []int{}
}
