/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    arr := make([]int, 0)
	var trav func (node *TreeNode)
	trav = func (node *TreeNode) {
		if node == nil {
			return
		}
		trav(node.Left)
		arr = append(arr, node.Val)
		trav(node.Right)
	}
	trav(root)
	return arr[k - 1]
}
