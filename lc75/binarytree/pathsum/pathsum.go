package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	cnt := 0

	if root.Val == targetSum {
		cnt++
	}

	arr := []int{root.Val}

	if root.Left != nil {
		cnt += traverse(root.Left, targetSum, arr)
	}

	if root.Right != nil {
		cnt += traverse(root.Right, targetSum, arr)
	}
	return cnt
}

func traverse(root *TreeNode, target int, arr []int) int {
	cnt := 0

	if root.Val == target {
		cnt++
	}

	var narr []int
	for _, v := range arr {
		narr = append(narr, v+root.Val)
		if v+root.Val == target {
			cnt++
		}
	}

	narr = append(narr, root.Val)

	if root.Left != nil {
		cnt += traverse(root.Left, target, narr)
	}

	if root.Right != nil {
		cnt += traverse(root.Right, target, narr)
	}

	return cnt
}

func pathSum_leet0ms(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, 0, map[int]int{})
}

func dfs(root *TreeNode, targetSum, prefixSum int, prefixMap map[int]int) int {
	if root == nil {
		return 0
	}
	count := 0

	prefixSum += root.Val
	if prefixSum == targetSum {
		count++
	}
	count += prefixMap[prefixSum-targetSum]

	prefixMap[prefixSum]++
	count += dfs(root.Left, targetSum, prefixSum, prefixMap)
	count += dfs(root.Right, targetSum, prefixSum, prefixMap)
	prefixMap[prefixSum]--
	return count
}
