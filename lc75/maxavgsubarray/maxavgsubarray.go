package longestones

func findMaxAverage(nums []int, k int) float64 {
	maxavg, l := -20000.0, len(nums)
	if l == 1 {
		return float64(nums[0])
	}

	avg, tk := 0.0, k
	for i, v := range nums {
		if tk != 0 {
			avg += float64(v)
			tk--
		} else {
			maxavg = max(maxavg, avg/float64(k))
			avg = avg - float64(nums[i-k]) + float64(v)
		}
	}

	maxavg = max(maxavg, avg/float64(k))

	return maxavg
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func findMaxAverage_leet138ms(nums []int, k int) float64 {
	// initial setup
	currSum := 0.00
	for i := 0; i < k; i++ {
		currSum += float64(nums[i])
	}

	var maxAverage float64

	maxAverage = currSum / float64(k)

	for i := k; i < len(nums); i++ {
		currSum += float64(nums[i])
		currSum -= float64(nums[i-k])

		if currAverage := currSum / float64(k); currAverage > maxAverage {
			maxAverage = currAverage
		}
	}

	return maxAverage
}
