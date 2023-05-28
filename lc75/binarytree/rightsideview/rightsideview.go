package rightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	arr := []int{}

	return traverse(root, arr, 1)
}

func traverse(root *TreeNode, arr []int, level int) []int {
	if root == nil {
		return arr
	}

	if len(arr) < level {
		arr = append(arr, root.Val)
	}

	arr = traverse(root.Right, arr, level+1)
	arr = traverse(root.Left, arr, level+1)

	return arr
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView_leet1ms(root *TreeNode) []int {
	var queue []*TreeNode
	var vals []int
	if root != nil {
		queue = append(queue, root)
		for len(queue) > 0 {
			queueSize := len(queue)
			var poppedVal *TreeNode
			for i := 0; i < queueSize; i++ {
				poppedVal = queue[0]
				queue = queue[1:]
				if poppedVal.Left != nil {
					queue = append(queue, poppedVal.Left)
				}
				if poppedVal.Right != nil {
					queue = append(queue, poppedVal.Right)
				}
			}
			vals = append(vals, poppedVal.Val)
		}
	}
	return vals
}
