/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    var dfs func (node *TreeNode, maxVal int) int
	dfs = func (node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}
		if node.Val >= maxVal {
			return dfs(node.Left, node.Val) + dfs(node.Right, node.Val) + 1
		} else {
			return dfs(node.Left, maxVal) + dfs(node.Right, maxVal)
		}
	}
	return dfs(root, math.MinInt)
}
