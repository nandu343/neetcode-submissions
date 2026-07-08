func searchMatrix(matrix [][]int, target int) bool {
    // l := [2]int{0,0}
    // r := [2]int{len(matrix), len(matrix[0])}
    // for l[0] <= r[0] && l[1] <= r[1] {
    //     mid := [2]int{(l[0]+r[0])/2, (l[1]+r[1])/2}
    //     if matrix[mid[0]][mid[1]] == target {
    //         return true
    //     }

    //     if matrix[mid[0]][mid[1]] >= target {
    //         r[0], r[1] = mid[0]
    //     }
    // }

    if len(matrix) * len(matrix[0]) == 1 {
        return matrix[0][0] == target
    }

    l := 0
    r := len(matrix) * len(matrix[0]) - 1
    for l <= r {
        m := (l+r)/2
        row := m / len(matrix[0])
        col := m % len(matrix[0])
        if matrix[row][col] == target {
            return true
        }

        if matrix[row][col] >= target {
            r = m - 1
        } else {
            l = m + 1
        }
    }
    return false
}
