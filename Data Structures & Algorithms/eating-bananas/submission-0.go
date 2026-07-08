func minEatingSpeed(piles []int, h int) int {
    var feas func (eat int) bool
    feas = func (eat int) bool {
        total := 0
        for _, v := range piles {
           total += v/eat
           if v % eat != 0 {
            total++
           } 
        }
        if total > h {
            return false
        }
        return true
    }
    r := 1000000000
    l := 1
    ans := r
    for l <= r {
        m := (r+l)/2
        if feas(m) {
            ans = m
            r = m - 1
        } else {
            l = m + 1
        }
    }
    return ans
}
