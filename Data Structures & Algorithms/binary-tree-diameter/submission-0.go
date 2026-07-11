/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    max := 0
	var dfs func (node *TreeNode) int 
	dfs = func (node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if (l + r) > max {
			max = l + r
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	dfs(root)
	return max
}
