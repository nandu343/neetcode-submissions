func findMin(arr []int) int {
	l, r := 0, len(arr)-1
    index := -1
    for l <= r {
        mid := l + (r-l)/2
        if arr[mid] <= arr[len(arr)-1] {
            index = mid
            r = mid - 1
            continue
        }
        l = mid + 1
    }
    return arr[index]
}
