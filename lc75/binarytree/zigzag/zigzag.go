package zigzag

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cnt int

func longestZigZagV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return traverseV2(root, 0, 0)
}

func traverseV2(root *TreeNode, left, right int) int {
	if root == nil {
		return 0
	}

	a := max(traverseV2(root.Left, right+1, 0), right)
	b := max(traverseV2(root.Right, 0, left+1), left)

	cnt := max(a, b)

	return cnt
}

func longestZigZag(root *TreeNode) int {
	cnt = 0
	if root == nil {
		return 0
	}

	traverse(root.Left, true)
	traverse(root.Right, false)

	return cnt
}

func traverse(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	a := traverse(root.Left, true)
	b := traverse(root.Right, false)

	if isLeft {
		b++
	} else {
		a++
	}

	cnt = max(cnt, a)
	cnt = max(cnt, b)

	if isLeft {
		return b
	}

	return a
}

func longestZigZag_leet90ms(root *TreeNode) int {
	path := new(int)
	var dfs func(*TreeNode, int, int, *int)
	dfs = func(root *TreeNode, left int, right int, path *int) {
		if root == nil {
			return
		}

		*path = max(*path, max(left, right))
		dfs(root.Left, right+1, 0, path)
		dfs(root.Right, 0, left+1, path)
	}

	dfs(root, 0, 0, path)
	return *path
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
