package validbst

import (
	"runtime"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var arr []int

func isValidBST(root *TreeNode) bool {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		funcName := details.Name()[len(details.Name())-10:]
		if funcName != "isValidBST" {
			arr = make([]int, 0)
		}
	}

	if root.Left == nil && root.Right == nil {
		arr = append(arr, root.Val)
		return true
	}

	if root.Left != nil {
		if !isValidBST(root.Left) {
			return false
		}
	}

	arr = append(arr, root.Val)

	if root.Right != nil {
		if !isValidBST(root.Right) {
			return false
		}
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}

	return true
}

// space O(N) because use recursion
// time O(N) because dfs
func isValidBST_leet0ms(root *TreeNode) bool {
	max := 1 << 31
	min := -1<<31 - 1

	return isValidBSTHelper(root, max, min)
}

func isValidBSTHelper(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return isValidBSTHelper(root.Left, root.Val, min) && isValidBSTHelper(root.Right, max, root.Val)

}
