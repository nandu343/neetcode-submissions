func missingNumber(nums []int) int {
	sum := 0
    sum2 := 0
    i := 0
    for _, v := range nums {
        sum += v
        sum2 += i
        i++
    }
    sum2 += i
    return sum2 - sum
}
