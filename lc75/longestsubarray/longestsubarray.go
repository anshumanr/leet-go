package longestsubarray

func longestSubarrayV2(nums []int) int {
	maxx, l := 0, len(nums)

	prevCnt, i, cnt := 0, 0, 0
	for ; i < l; i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			// if prevCnt+cnt > maxx {
			// 	maxx = prevCnt + cnt
			// }
			maxx = max(maxx, prevCnt+cnt)
			prevCnt = cnt
			cnt = 0
		}
	}

	// if prevCnt+cnt > maxx {
	// 	maxx = prevCnt + cnt
	// }
	maxx = max(maxx, prevCnt+cnt)

	if maxx == l {
		return maxx - 1
	}

	return maxx
}

/*
 * Using the same algo as finding longest 1s with k replacements
 */
func longestSubarray(nums []int) int {
	max, l := 0, len(nums)

	zi, tk, i, cnt := -1, 1, 0, 0
	for ; i < l; i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			if tk == 0 {
				t, oc := zi+1, 0
				for ; t < i; t++ {
					oc++
					if nums[t] == 0 {
						break
					}
				}
				cnt = oc + 1
				zi = t
			} else {
				tk--
				cnt++
				if zi < 0 {
					zi = i
				}
			}
		}

		if cnt > max {
			max = cnt
		}
	}

	return max - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestSubarray_leet25ms(nums []int) int {

	prevCnt, cnt, maxSize := 0, 0, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			maxSize = max(maxSize, cnt+prevCnt)
			prevCnt = cnt
			cnt = 0
		}
	}

	maxSize = max(maxSize, cnt+prevCnt)

	if maxSize == len(nums) {
		return maxSize - 1
	} else {
		return maxSize
	}

}
