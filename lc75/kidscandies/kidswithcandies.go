package kidscandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	arr := make([]bool, n)
	max := 0
	for i := 0; i < n; i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= max {
			arr[i] = true
		}
	}
	return arr
}

func kidsWithCandies_leet1ms(candies []int, extraCandies int) []bool {
	n := len(candies)
	ans := make([]bool, n)
	max := 0
	for i := 0; i < n; i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := 0; i < n; i++ {
		ans[i] = candies[i]+extraCandies >= max
	}
	return ans
}
