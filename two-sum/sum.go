package twosum

func twoSumAttempt1(nums []int, target int) []int {
	var result []int
	result = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j

				return result
			}
		}
	}

	return result
}

func twoSumAttempt2(nums []int, target int) []int {
	var result []int

	for i, v1 := range nums {
		x := target - v1
		for j, v2 := range nums[i+1:] {

			if v2 == x {
				return []int{i, j}
			}
		}
	}

	return result
}

func twoSumMapInt32ms(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return append(result, i, j)
			}
		}
	}

	return result
}

func twoSumAttempt3(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

func twoSumMapInt8ms(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if idx, ok := m[v]; ok {
			return []int{idx, i}
		}
		m[target-v] = i
	}
	return []int{-1, -1}
}

func twoSumMapInt12ms(nums []int, target int) []int {
	var table = make(map[int]int)
	var result [2]int
	for i, num := range nums {
		if val, ok := table[target-num]; ok {
			result[0] = val
			result[1] = i
		} else {
			table[num] = i
		}
	}
	return result[:2]
}

func twoSumMapInt4ms(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, val := range nums {
		comp := target - val
		index, ok := lookup[comp]
		if ok {
			return []int{i, index}
		}
		lookup[val] = i
	}
	return []int{}
}

func twoSumMapInt0ms(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, val := range nums {
		j, ok := numsMap[target-val]
		if ok && i != j {
			return []int{i, j}
		}
		numsMap[val] = i
	}
	return nil
}

func twoSumMapInt20ms(nums []int, target int) []int {
	for i, x := range nums {
		j := indexOf(nums[i+1:], target-x)
		if j >= 0 {
			return []int{i, j + i + 1}
		}
	}
	return []int{}
}

func indexOf(nums []int, v int) int {
	for i, x := range nums {
		if x == v {
			return i
		}
	}
	return -1
}
