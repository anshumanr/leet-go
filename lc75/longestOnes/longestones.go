package longestones

func longestOnes(nums []int, k int) int {
	max, l := 0, len(nums)

	i, gap, cnt, oc := 0, 0, 0, 0
	for ; i < l; i++ {
		if nums[i] == 0 {
			gap++
			if gap <= k {
				cnt++
			} else {
				oc = findw(nums[i-cnt : i])
				cnt -= oc
				gap = k
			}
		} else {
			cnt++
		}
		if cnt > max {
			max = cnt
		}

	}

	return max
}

func findw(arr []int) int {
	oc := 0
	for _, n := range arr {
		if n == 0 {
			return oc
		}
		oc++
	}

	return oc
}

func longestOnes_leet47ms(nums []int, k int) int {
	zero_pointer := 0
	count := 0
	max := 0
	for idx, _ := range nums {
		if nums[idx] == 0 {
			if k > 0 {
				k--
			} else {

				// move zero_pointer till you skip atleast one zero.
				for zero_pointer < idx+1 {
					count--
					if nums[zero_pointer] == 0 {
						zero_pointer++
						break
					}
					zero_pointer++
				}
			}
		}

		count++

		if count > max {
			max = count
		}
	}

	return max
}
