package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
  Given the root of a binary tree, return the level order
  traversal of its nodes' values. (i.e., from left to right, level by level).
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var levelorder [][]int

	queue := make([]*TreeNode, 2000)
	qf, qb := 0, 0
	queue[qb] = root
	qb++
	for qf < qb {
		lvb := qb
		level := []int{}
		for qf < lvb {
			curr := queue[qf]
			qf++

			level = append(level, curr.Val)

			if curr.Left != nil {
				queue[qb] = curr.Left
				qb++
			}
			if curr.Right != nil {
				queue[qb] = curr.Right
				qb++
			}
		}
		levelorder = append(levelorder, level)
	}

	return levelorder
}

func levelOrder_leet1ms(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	q := []*TreeNode{root}

	for len(q) != 0 {
		var v []int
		var _q []*TreeNode
		for _, n := range q {
			v = append(v, n.Val)
			if n.Left != nil {
				_q = append(_q, n.Left)
			}
			if n.Right != nil {
				_q = append(_q, n.Right)
			}
		}

		q = _q
		result = append(result, v)
	}

	return result
}
