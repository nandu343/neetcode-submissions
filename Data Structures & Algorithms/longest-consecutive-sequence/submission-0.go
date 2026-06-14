func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
    for _, v := range nums {
        set[v] = struct{}{}
    }

    if len(nums) == 0 {
        return 0
    }

    long := 0
    for k, _ := range set {
        v := k - 1
        count := 0
        for {
            if _, e := set[v]; e == false {
                break
            }
            count++
            v--
        }
        if count > long {
            long = count
        }
    }
    
    return long + 1
}
