func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
    postfix := make([]int, len(nums))
    prefix[0] = 1
    postfix[len(nums) - 1] = 1
    for i := 0; i < len(nums) - 1; i++ {
        prefix[i + 1] = prefix[i] * nums[i]
        postfix[len(nums)-2 - i] = postfix[len(nums)-1 - i] * nums[len(nums)-1-i]
    }

    for i := 0; i < len(nums) - 1; i++ {
        prefix[i] = prefix[i] * postfix[i]
    }
    return prefix
}
