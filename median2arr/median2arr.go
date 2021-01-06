package median2arr

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	var arr [2000]int

	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			arr[k] = nums1[i]
			i++
		} else {
			arr[k] = nums2[j]
			j++
		}
		k++
	}

	for ; i < l1; i++ {
		arr[k] = nums1[i]
		k++
	}
	for ; j < l2; j++ {
		arr[k] = nums2[j]
		k++
	}

	if (l+1)%2 == 0 {
		return (float64)(arr[(l+1)/2-1])
	} else {
		h := l / 2
		return ((float64)(arr[h-1]+arr[h]) / 2)
	}
}

func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	arr := make([]int, l)

	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			arr[k] = nums1[i]
			i++
		} else {
			arr[k] = nums2[j]
			j++
		}
		k++
	}

	for ; i < l1; i++ {
		arr[k] = nums1[i]
		k++
	}
	for ; j < l2; j++ {
		arr[k] = nums2[j]
		k++
	}

	if (l+1)%2 == 0 {
		return (float64)(arr[(l+1)/2-1])
	} else {
		h := l / 2
		return ((float64)(arr[h-1]+arr[h]) / 2)
	}
}

func findMedianSortedArraysV3(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	mid1, mid2 := -1, -1
	if l%2 == 0 {
		mid1 = l / 2
		mid2 = mid1 - 1
	} else {
		mid1 = (l - 1) / 2
	}

	h1, h2 := 0, 0
	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			if k == mid1 {
				h1 = nums1[i]
				return median(mid1, mid2, h1, h2)
			}
			if k == mid2 {
				h2 = nums1[i]
			}
			i++
		} else {
			if k == mid1 {
				h1 = nums2[j]
				return median(mid1, mid2, h1, h2)
			}
			if k == mid2 {
				h2 = nums2[j]
			}
			j++
		}
		k++
	}

	for ; i < l1; i++ {
		if k == mid1 {
			h1 = nums1[i]
			return median(mid1, mid2, h1, h2)
		}
		if k == mid2 {
			h2 = nums1[i]
		}
		k++
	}
	for ; j < l2; j++ {
		if k == mid1 {
			h1 = nums2[j]
			return median(mid1, mid2, h1, h2)
		}
		if k == mid2 {
			h2 = nums2[j]
		}
		k++
	}

	return median(mid1, mid2, h1, h2)
}

func findMedianSortedArraysV4(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	mid1, mid2 := -1, -1
	if l%2 == 0 {
		mid1 = l / 2
		mid2 = mid1 - 1
	} else {
		mid1 = (l - 1) / 2
	}

	h1, h2 := 0, 0
	i, j, k := 0, 0, 0
	oldi, oldj := 0, 0
	for i < l1 && j < l2 && i+j < mid1-1 {
		k = i + j
		if nums1[i] <= nums2[j] {
			if i == 0 {
				i = 1
				k++
			} else {
				oldi = i
				i *= 2
			}
		} else {
			if j == 0 {
				j = 1
				k++
			} else {
				oldj = j
				j *= 2
			}
		}
	}

	i, j = oldi, oldj
	k = i + j
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			if k == mid1 {
				h1 = nums1[i]
				return median(mid1, mid2, h1, h2)
			}
			if k == mid2 {
				h2 = nums1[i]
			}
			i++
		} else {
			if k == mid1 {
				h1 = nums2[j]
				return median(mid1, mid2, h1, h2)
			}
			if k == mid2 {
				h2 = nums2[j]
			}
			j++
		}
		k++
	}

	for ; i < l1; i++ {
		if k == mid1 {
			h1 = nums1[i]
			return median(mid1, mid2, h1, h2)
		}
		if k == mid2 {
			h2 = nums1[i]
		}
		k++
	}
	for ; j < l2; j++ {
		if k == mid1 {
			h1 = nums2[j]
			return median(mid1, mid2, h1, h2)
		}
		if k == mid2 {
			h2 = nums2[j]
		}
		k++
	}

	return median(mid1, mid2, h1, h2)
}

func findMedianSortedArraysV5(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2

	mid1, mid2 := -1, -1
	if l%2 == 0 {
		mid1 = l / 2
		mid2 = mid1 - 1
	} else {
		mid1 = (l - 1) / 2
	}

	h1, h2 := 0, 0
	i, j, k := 0, 0, 0
	oldi, oldj := 0, 0
	for i < l1 && j < l2 && i+j < mid1-1 {
		k = i + j
		if nums1[i] <= nums2[j] {
			if i == 0 {
				i = 1
				k++
			} else {
				oldi = i
				i *= 2
			}
		} else {
			if j == 0 {
				j = 1
				k++
			} else {
				oldj = j
				j *= 2
			}
		}
	}

	i, j = oldi, oldj
	k = i + j
	for i < l1 && j < l2 {
		if nums1[i] <= nums2[j] {
			if k == mid1 {
				h1 = nums1[i]
				return median(mid1, mid2, h1, h2)
			}
			if k == mid2 {
				h2 = nums1[i]
			}
			i++
		} else {
			if k == mid1 {
				h1 = nums2[j]
				return median(mid1, mid2, h1, h2)
			}
			if k == mid2 {
				h2 = nums2[j]
			}
			j++
		}
		k++
	}

	for ; i < l1; i++ {
		if k == mid1 {
			h1 = nums1[i]
			return median(mid1, mid2, h1, h2)
		}
		if k == mid2 {
			h2 = nums1[i]
		}
		k++
	}
	for ; j < l2; j++ {
		if k == mid1 {
			h1 = nums2[j]
			return median(mid1, mid2, h1, h2)
		}
		if k == mid2 {
			h2 = nums2[j]
		}
		k++
	}

	return median(mid1, mid2, h1, h2)
}

func median(mid1, mid2, h1, h2 int) float64 {
	if mid2 < 0 {
		return (float64)(h1)
	} else {
		return ((float64)(h1+h2) / 2)
	}
}

func findMedianSortedArrays4ms(nums1 []int, nums2 []int) float64 {
	var median int

	oneArrayEmpty := func(a []int) float64 {
		median = a[(len(a) / 2)]
		finalMedian := float64(median)
		if len(a)%2 == 0 {
			median2 := a[(len(a)/2)-1]
			finalMedian = (finalMedian + float64(median2)) / 2
		}
		return finalMedian
	}
	if len(nums1) == 0 {
		return oneArrayEmpty(nums2)
	}
	if len(nums2) == 0 {
		return oneArrayEmpty(nums1)
	}

	targetPosition := ((len(nums1) + len(nums2)) / 2)
	median = optimalIndexPosition(0, targetPosition, nums1, nums2)
	if median == -1 {
		median = optimalIndexPosition(0, targetPosition, nums2, nums1)
	}
	finalMedian := float64(median)
	if (len(nums1)+len(nums2))%2 == 0 {
		secondMedian := optimalIndexPosition(0, targetPosition-1, nums1, nums2)
		if secondMedian == -1 {
			secondMedian = optimalIndexPosition(0, targetPosition-1, nums2, nums1)
		}
		finalMedian = (finalMedian + float64(secondMedian)) / 2
	}
	return finalMedian
}

func optimalIndexPosition(startIndex int, targetPos int, a []int, b []int) int {

	if len(a) < 1 {
		return -1
	}
	if targetPos == 0 && startIndex == 0 {
		if a[0] <= b[0] {
			return a[targetPos]
		}
	}
	medianPos := len(a) / 2
	posInOtherArray := posInArray(0, a[medianPos], b)
	idealPositon := startIndex + medianPos + posInOtherArray
	if len(a) == 1 {
		if targetPos == idealPositon {
			return a[medianPos]
		}
		return -1
	}
	if idealPositon == targetPos {
		return a[medianPos]
	}
	if idealPositon < targetPos {
		startIndex += medianPos
		return optimalIndexPosition(startIndex, targetPos, a[medianPos:], b)
	}
	return optimalIndexPosition(startIndex, targetPos, a[:medianPos], b)
}

func posInArray(startindex int, element int, nums []int) int {

	if len(nums) == 1 {
		if element < nums[0] {
			return 0 + startindex
		}
		return 1 + startindex
	}
	if len(nums) == 0 {
		return 0
	}
	medianPos := len(nums) / 2

	if element >= nums[medianPos-1] && element <= nums[medianPos] {
		return startindex + medianPos
	}
	if element < nums[medianPos-1] {
		return posInArray(startindex, element, nums[:medianPos])
	}
	startindex += medianPos
	return posInArray(startindex, element, nums[medianPos:])
}

func findMedianSortedArrays8ms(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	sorted := make([]int, 0, l1+l2)
	var (
		i int
		j int
	)
	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	if i < l1 {
		sorted = append(sorted, nums1[i:]...)
	}
	if j < l2 {
		sorted = append(sorted, nums2[j:]...)
	}
	l := len(sorted)
	if l%2 == 1 {
		return float64(sorted[l/2])
	}
	low := sorted[l/2]
	high := sorted[l/2-1]
	return float64(low+high) / 2

}
