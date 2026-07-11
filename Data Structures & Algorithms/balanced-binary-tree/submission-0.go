/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var dfs func (node *TreeNode) (int, bool)
	dfs = func (node *TreeNode) (int, bool) {
		if node == nil {
			return -1, true
		}
		l, a := dfs(node.Left)
		r, b := dfs(node.Right)
		if !a || !b {
			return 0, false
		}
		if l - r > 1 || l - r < -1{
			return 0, false
		}
		if l > r {
			return l + 1, true
		} else {
			return r + 1, true
		}
	}
	_, ans := dfs(root)
	return ans
}
