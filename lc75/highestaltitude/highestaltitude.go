package highestaltitude

func largestAltitude(gain []int) int {
	max, l := 0, len(gain)

	i, cnt := 0, 0
	for ; i < l; i++ {
		cnt += gain[i]
		if cnt > max {
			max = cnt
		}
	}

	return max
}

func largestAltitude_leet0ms(gain []int) int {
	ans := max(gain[0], 0)
	for i := 1; i < len(gain); i++ {
		gain[i] += gain[i-1]
		ans = max(ans, gain[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
