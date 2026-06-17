func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    money := make([]int, len(nums) - 1)
    money2 := make([]int, len(nums) - 1)
    var robber func (nums []int, arr []int) int
    robber = func (nums []int, arr []int) int {
        if len(nums) == 1 {
            return nums[0]
        } else if len(nums) == 2 {
            return max(nums[0],nums[1])
        }
        arr[0] = nums[0]
        arr[1] = max(arr[0], nums[1])
        for i := 2; i < len(nums); i++ {
            arr[i] = max(arr[i-1], arr[i-2]+nums[i])
        }
        return arr[len(nums)-1]
    }
    a := robber(nums[:len(nums)-1], money)
    b := robber(nums[1:], money2)
    return max(a,b)
}


func max (a int, b int) int {
    if a > b {
        return a
    }
    return b
}
