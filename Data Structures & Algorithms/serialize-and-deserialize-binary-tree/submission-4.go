/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    
	if root == nil {
		return ""
	}
	
	
	queue := []*TreeNode{root}
	res := strconv.Itoa(root.Val)
	ans := []string{res}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil {
			ans = append(ans, "nil")
		} else {
			res2:= strconv.Itoa(node.Left.Val)
			ans = append(ans, res2)
			queue = append(queue, node.Left)
		}

		if node.Right == nil {
			ans = append(ans, "nil")
		} else {
			res2 := strconv.Itoa(node.Right.Val)
			ans = append(ans, res2)
			queue = append(queue, node.Right)
		}

	}
	return strings.Join(ans, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// Handle the empty/nil tree edge case
	if data == "" || data == "nil" {
		return nil
	}

	// Split the serialization string into tokens
	arr := strings.Split(data, " ")

	// Rebuild the root node
	val, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: val}

	// Use a queue to track nodes whose children we need to assign
	queue := []*TreeNode{root}
	index := 1 // Start reading tokens from index 1

	for len(queue) != 0 && index < len(arr) {
		parent := queue[0]
		queue = queue[1:]

		// 1. Process Left Child
		if arr[index] != "nil" {
			leftVal, _ := strconv.Atoi(arr[index])
			parent.Left = &TreeNode{Val: leftVal}
			queue = append(queue, parent.Left) // Push to queue to find its children later
		}
		index++ // Move to the next token

		// Guard check in case the array ends abruptly
		if index >= len(arr) {
			break
		}

		// 2. Process Right Child
		if arr[index] != "nil" {
			rightVal, _ := strconv.Atoi(arr[index])
			parent.Right = &TreeNode{Val: rightVal}
			queue = append(queue, parent.Right) // Push to queue to find its children later
		}
		index++ // Move to the next token
	}

	return root
}
