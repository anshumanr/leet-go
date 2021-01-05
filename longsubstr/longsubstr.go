package longsubstr

type btree struct {
	b     byte
	v     bool
	index int
	left  *btree
	right *btree
}

func lengthOfLongestSubstringV1(s string) int {
	arr := btree{}
	max := 0
	cnt := 0

	l := len(s)
	for i := 0; i < l; i++ {
		f, ix := add(&arr, s[i], i)
		if f == true {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
			arr = btree{}
			i = ix
		}
	}

	if cnt > max {
		max = cnt
	}

	//fmt.Println("longest len: ", max)

	return max
}

func add(arr *btree, c byte, ix int) (bool, int) {
	a := arr

	for a.v != false {
		if a.b == c {
			return false, a.index
		} else if c < a.b {
			if a.left == nil {
				a.left = &btree{b: c, v: true, index: ix}
				//fmt.Println("adding left: ", c)
				return true, a.index
			} else {
				a = a.left
			}
		} else {
			if a.right == nil {
				a.right = &btree{b: c, v: true, index: ix}
				//fmt.Println("adding right: ", c)
				return true, a.index
			} else {
				a = a.right
			}
		}
	}

	//not found
	if arr.v == false {
		arr.v, arr.b, arr.index = true, c, ix
		//fmt.Println("adding to root: ", c)
	}

	return true, arr.index
}

func lengthOfLongestSubstringV2(s string) int {
	arr := make([]int, 128)
	ind := make([]int, 128)

	max, low, cnt := 0, 0, 0

	l := len(s)
	for i := 0; i < l; i++ {
		t := (int)(s[i])
		arr[t]++
		if arr[t] > 1 {
			if cnt > max {
				max = cnt
			}

			if ind[t] > low {
				low = ind[t]
			}
			cnt = i - low
			ind[t] = i
			arr[t]--
		} else {
			ind[t] = i
			cnt++
		}
	}

	if cnt > max {
		max = cnt
	}

	//fmt.Println("longest len: ", max)

	return max
}

func lengthOfLongestSubstring0ms(s string) int {
	m := make([]int, 128)
	fast, slow := 0, 0
	longest := 0

	for fast < len(s) {
		m[int(s[fast])]++
		for m[int(s[fast])] > 1 {
			m[int(s[slow])]--
			slow++
		}
		fast++
		if fast-slow > longest {
			longest = fast - slow
		}
	}

	return longest
}
