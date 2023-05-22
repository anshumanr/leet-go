package goodnodes

/*
 * Given a binary tree root, a node X in the tree is named good if in the
 * path from root to X there are no nodes with a value greater than X.
 * Return the number of good nodes in the binary tree.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return traverse(root, root.Val)
}

func traverse(root *TreeNode, highVal int) int {
	cnt := 0
	if root.Val >= highVal {
		cnt++
		highVal = root.Val
	}

	if root.Left != nil {
		cnt += traverse(root.Left, highVal)
	}
	if root.Right != nil {
		cnt += traverse(root.Right, highVal)
	}

	return cnt
}

func dfs(root *TreeNode, maxParent int) int {
	if root == nil {
		return 0
	}
	count := 0
	if maxParent <= root.Val {
		count++
		maxParent = root.Val
	}
	return dfs(root.Left, maxParent) + dfs(root.Right, maxParent) + count
}

func goodNodes_leet65ms(root *TreeNode) int {
	return dfs(root, root.Val-1)
}
