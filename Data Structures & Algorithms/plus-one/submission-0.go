func plusOne(digits []int) []int {
    carry := 1
	for i := len(digits) - 1; i > -1; i-- {
		val := digits[i] + carry
		if val >= 10 {
			val %= 10
			carry = 1
		} else {
			carry = 0
		}
		digits[i] = val
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
