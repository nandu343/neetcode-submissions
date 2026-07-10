/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return nil
	}
	ans := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if i == 0 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}
