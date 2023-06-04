package deletenode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	return traverse(root, nil, false, key)
}

func traverse(root, parent *TreeNode, isLeft bool, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Right == nil {
			if parent != nil {
				if isLeft {
					parent.Left = root.Left
				} else {
					parent.Right = root.Left
				}
			} else {
				root = root.Left
			}
			return root
		}
		if root.Left == nil {
			if parent != nil {
				if isLeft {
					parent.Left = root.Right
				} else {
					parent.Right = root.Right
				}
			} else {
				root = root.Right
			}
			return root
		}

		succ := promote(root.Right, root)
		if succ == null && parent == nil {
			return nil
		}

		root.Val = succ

		return root
	}

	if root.Val > key {
		traverse(root.Left, root, true, key)
	} else {
		traverse(root.Right, root, false, key)
	}

	return root
}

func promote(root, parent *TreeNode) int {
	if root == nil {
		return null
	}

	if root.Left != nil {
		return promote(root.Left, root)
	}

	if parent != nil {
		if parent.Left == root {
			parent.Left = root.Right
		} else {
			parent.Right = root.Right
		}
	}

	return root.Val
}

func deleteNode_leet0ms(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = findSuccessor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = findPredecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func findPredecessor(root *TreeNode) int {
	curr := root.Left
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Val
}

func findSuccessor(root *TreeNode) int {
	curr := root.Right
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Val
}
