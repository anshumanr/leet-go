package mostwater

func maxArea(height []int) int {
	max, l := 0, len(height)

	i, j := 0, l-1
	for i < j {
		m := 0
		if height[i] <= height[j] {
			m = height[i] * (j - i)
			i++
		} else {
			m = height[j] * (j - i)
			j--
		}
		if m > max {
			max = m
		}
	}

	return max
}

func maxArea_leet6ms(height []int) int {
	n := len(height)
	l, r := 0, n-1

	result := 0
	for l <= r {
		area := (r - l) * min(height[l], height[r])
		result = max(result, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea_BF(height []int) int {
	max, l := 0, len(height)

	i := 0
	for ; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			m := minn(height[i], height[j])
			if max < (m * (j - i)) {
				max = m * (j - i)
			}
		}
	}
	return max
}

func minn(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
