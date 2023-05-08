package maxksumpair

import "sort"

func maxOperations(nums []int, k int) int {
	cnt, l := 0, len(nums)

	sort.Ints(nums)

	i, j := 0, l-1
	for i < j {
		s := nums[i] + nums[j]
		if s == k {
			cnt++
			i++
			j--
		} else if s < k {
			i++
		} else {
			//s > k
			j--
		}
	}

	return cnt
}

func maxOperations_leet96ms(nums []int, k int) int {
	valuesMap := make(map[int]int)
	count := 0
	for _, val := range nums {
		if val < k {
			if c, ok := valuesMap[k-val]; ok && c != 0 {
				count++
				valuesMap[k-val] = valuesMap[k-val] - 1
				continue
			}

			valuesMap[val] = valuesMap[val] + 1
		}
	}

	return count
}

func maxOperationsV2(nums []int, k int) int {
	cnt, l := 0, len(nums)
	stack, top := make([]int, l), 0

	i, j := 0, 0
	for i < l-1 {
		if nums[i] == 0 || nums[i] >= k {
			i++
			continue
		}
		j = i + 1
		t := i
		for j < l {
			if nums[j] == 0 {
				j++
				continue
			}
			s := nums[t] + nums[j]
			if s == k {
				nums[t], nums[j] = 0, 0
				cnt++
				if top > 0 && t != stack[top-1] {
					top--
					t = stack[top]
				} else {
					break
				}
			} else if nums[t] == nums[j] {
				stack[top] = j
				top++
			}
			j++
		}
		top = 0
		i++
	}

	return cnt
}

func maxOperations_BF(nums []int, k int) int {
	cnt, l := 0, len(nums)

	i, j := 0, 0
	for i < l-1 {
		if nums[i] == 0 {
			i++
			continue
		}
		j = i + 1
		for j < l {
			if nums[j] == 0 {
				j++
				continue
			}
			s := nums[i] + nums[j]
			if s == k {
				nums[i], nums[j] = 0, 0
				cnt++
				break
			}
			j++
		}
		i++
	}

	return cnt
}
