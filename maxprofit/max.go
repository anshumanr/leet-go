package maxprofit

import (
	"math"
)

func maxProfit(nums []int) int {
	lowI, highI, buy, sell := 0, 0, nums[0], nums[0]

	profit := 0
	for i, v := range nums {
		if v < buy {
			buy, lowI = v, i
			sell, highI = v, i
		}
		if v >= sell {
			sell = v
			highI = i

			p := nums[highI] - nums[lowI]

			if p > profit {
				profit = p
			}
		}

	}

	return profit
}

func maxProfit_leet99ms(prices []int) int {
	maxProfit := 0
	buy := math.MaxUint32
	for i := 0; i < len(prices); i++ {
		if prices[i] > buy {
			if prices[i]-buy > maxProfit {
				maxProfit = prices[i] - buy
			}
		} else {
			buy = prices[i]
		}
	}
	return maxProfit
}
