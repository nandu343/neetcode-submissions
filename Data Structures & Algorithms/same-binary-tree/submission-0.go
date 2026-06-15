/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func (a *TreeNode, b *TreeNode) bool
	dfs = func (a *TreeNode, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}
		
		if a.Val != b.Val {
			return false
		}

		return dfs(a.Left, b.Left) && dfs(a.Right, b.Right)
	}

	return dfs(p, q)
    
}
