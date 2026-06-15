/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    var dfs func (node *TreeNode) int
	dfs = func (node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left) + 1
		r := dfs(node.Right) + 1
		if l > r {
			return l
		} else {
			return r
		}
	}
	return dfs(root)
}
