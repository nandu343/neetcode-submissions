import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    ans := make([][]int, 0)
    
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i - 1] {
            continue
        }
        
        target := 0 - nums[i]
        l := i + 1
        r := len(nums) - 1
        
        for l < r {
            currentSum := nums[l] + nums[r]
            
            if currentSum == target {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l++
                r--
                
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            } else if currentSum > target {
                r--
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            } else {
                l++
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            }
        }
    }
    return ans
}