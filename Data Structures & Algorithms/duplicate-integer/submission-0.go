func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{})

    for _, v := range nums {
        if _, e := set[v]; e == true {
            return true
        }
        set[v] = struct{}{}
    }
    return false
}
