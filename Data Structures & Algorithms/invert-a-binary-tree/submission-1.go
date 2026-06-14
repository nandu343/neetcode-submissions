/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(tree *TreeNode) *TreeNode {
    var dfs func (node *TreeNode) 
    dfs = func (node *TreeNode) {
        if node == nil {
            return
        }

        node.Left, node.Right = node.Right, node.Left
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(tree)
    return tree
}
