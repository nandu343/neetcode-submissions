/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    var dfs func (node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode
	dfs = func (node *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		if p.Val <= node.Val && q.Val >= node.Val || p.Val >= node.Val && q.Val <= node.Val {
			return node
		} else if p.Val < node.Val && q.Val < node.Val {
			return dfs(node.Left, p, q)
		} else {
			return dfs(node.Right, p , q)
		}
	}
	return dfs(root, p, q)
}
