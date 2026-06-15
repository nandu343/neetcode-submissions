/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var dfs func (a *TreeNode, b *TreeNode) bool
	dfs = func (a *TreeNode, b *TreeNode) bool {
		if a == nil || b == nil {
			return false
		}

		return checkSubtree(a, b) || isSubtree(a.Left, b) || isSubtree(a.Right, b)
	}
	return dfs(root, subRoot)
}

func checkSubtree (root *TreeNode, check *TreeNode) bool {
	if root == nil && check == nil {
		return true
	}
	if root == nil || check == nil {
		return false
	}

	if root.Val == check.Val {
		return true && checkSubtree(root.Left, check.Left) && checkSubtree(root.Right, check.Right)
	}
	return false
}
