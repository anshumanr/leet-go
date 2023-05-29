package maxlevelsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	var lvlSum []int
	lvlSum = traverse(root, 0, lvlSum)

	maxSum, lvl := -1000000, 1
	for i, v := range lvlSum {
		if v > maxSum {
			maxSum = v
			lvl = i + 1
		}
	}

	return lvl
}

func traverse(root *TreeNode, lvl int, lvlsum []int) []int {
	if root == nil {
		return lvlsum
	}

	if len(lvlsum) < lvl+1 {
		lvlsum = append(lvlsum, 0)
	}

	lvlsum[lvl] += root.Val
	lvlsum = traverse(root.Left, lvl+1, lvlsum)
	lvlsum = traverse(root.Right, lvl+1, lvlsum)

	return lvlsum
}

func maxLevelSumV1(root *TreeNode) int {
	maxSum := -1000000
	q := []*TreeNode{root}

	lvl, maxlvl := 1, 1
	for len(q) > 0 {
		sum := 0
		for _, k := range q {
			sum += k.Val
		}
		if sum > maxSum {
			maxSum = sum
			maxlvl = lvl
		}

		q2 := []*TreeNode{}
		for _, k := range q {
			if k.Left != nil {
				q2 = append(q2, k.Left)
			}
			if k.Right != nil {
				q2 = append(q2, k.Right)
			}
		}
		lvl++
		q = q2
	}

	return maxlvl
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum_leet100ms(root *TreeNode) int {
	max := root.Val
	cur := 1
	ans := 1
	lvl := [](*TreeNode){root}
	for len(lvl) > 0 {
		lvl2 := [](*TreeNode){}
		sum := 0
		for _, n := range lvl {
			sum += n.Val
			if n.Left != nil {
				lvl2 = append(lvl2, n.Left)
			}
			if n.Right != nil {
				lvl2 = append(lvl2, n.Right)
			}
		}
		// fmt.Println(cur, sum)
		if sum > max {
			max = sum
			ans = cur
		}
		cur++
		lvl = lvl2
	}
	return ans
}
