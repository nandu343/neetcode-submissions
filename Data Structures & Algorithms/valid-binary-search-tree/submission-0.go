/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var dfs func (node *TreeNode, min int, max int) bool
	dfs = func (node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		if node.Val > min && node.Val < max {
			return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
		}
		return false
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
