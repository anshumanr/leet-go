package lca

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}

	panc := lowestCommonAncestor(root.Left, p, q)
	qanc := lowestCommonAncestor(root.Right, p, q)

	if panc == nil {
		return qanc
	}
	if qanc == nil {
		return panc
	}

	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor_leet0ms(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}
