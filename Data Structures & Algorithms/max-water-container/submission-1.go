func maxArea(arr []int) int {
	l, r := 0, len(arr) - 1
    max := -1
    for l < r {
        s1 := min(arr[l], arr[r])
        s2 := r - l
        area := s1 * s2
        if area > max {
            max = area
        }
        
        if min(arr[l], arr[r]) == arr[l] {
            l++
        } else {
            r--
        }
    }
    return max
}
