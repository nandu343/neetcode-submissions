func maxPathSum(root *TreeNode) int {
    maxa := root.Val
    
    var dfs func (node *TreeNode) int
    dfs = func (node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        lsum := dfs(node.Left)
        rsum := dfs(node.Right)

        // 1. Update the global max (maxa).
        // The best path at THIS node could be:
        // just the node itself, node+left, node+right, or node+left+right (the arch)
        maxa = max(maxa, node.Val)
        maxa = max(maxa, node.Val + lsum)
        maxa = max(maxa, node.Val + rsum)
        maxa = max(maxa, node.Val + lsum + rsum)
        
        // 2. Return ONLY a valid straight path to the parent.
        // The parent can only accept: just this node, or this node + ONE child.
        return max(node.Val, node.Val + max(lsum, rsum))
    }

    dfs(root)
    return maxa
}