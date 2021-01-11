package longpalin

func longestPalindromeV1(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	if l == 2 {
		if s[0] != s[1] {
			return string(s[0])
		} else {
			return s
		}
	}

	subs := ""
	left, right := -1, -1
	i, k := 0, l-1
	for ; i < l && i <= k; i++ {
		if s[i] == s[k] {
			if left == -1 || right == -1 {
				left, right = i, k
			}
		} else {
			left, right = -1, -1
		}
		k--
	}

	if left != -1 && right != -1 {
		subs = s[left : right+1]
		if l == len(subs) {
			return s
		}
	}

	sub1 := longestPalindromeV1(s[:l-1])
	sub2 := longestPalindromeV1(s[1:])
	if len(sub1) > len(subs) {
		subs = sub1
	}
	if len(sub2) > len(subs) {
		subs = sub2
	}

	return subs
}

func longestPalindromeV2(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	if l == 2 {
		if s[0] != s[1] {
			return string(s[0])
		} else {
			return s
		}
	}

	plen := l
	subs := ""

	for plen > 0 {
		start, size := 0, plen-1

		if len(subs) >= plen {
			return subs
		}

		for start+size < l {
			pal := ""
			for i := start; i+size < l; i++ {
				pal = findPalindrome(s[i : i+size+1])
				if len(pal) > len(subs) {
					subs = pal
				}
			}

			start++
		}

		plen--
	}

	return subs
}

func longestPalindromeV3(s string) string {
	l := len(s)
	subs := ""

	switch l {
	case 0:
		return ""
	case 1:
		return s
	case 2:
		if s[0] != s[1] {
			return string(s[0])
		} else {
			return s
		}
		// case 3:
		// 	if s[0] == s[2] {
		// 		return s
		// 	} else if s[0] == s[1] {
		// 		return s[:2]
		// 	} else if s[1] == s[2] {
		// 		return s[1:]
		// 	} else {
		// 		return string(s[0])
		// 	}
	}

	epvt := 1
	opvt := 2
	even, odd := false, false
	i := 0
	for i = 0; i < l; i++ {
		if i == 0 {
			subs = string(s[i])
			continue
		}

		//even length palindrome
		if i-epvt >= 0 && s[i] == s[i-epvt] {
			epvt += 2
			even = true
		} else if even {
			even = false
			t := s[i-epvt+1 : i]
			epvt = 1
			if len(t) > len(subs) {
				subs = t
			}
		}

		//odd length palindrome
		if i-opvt >= 0 && s[i] == s[i-opvt] {
			opvt += 2
			odd = true
		} else if odd {
			odd = false
			t := s[i-opvt+1 : i]
			opvt = 2
			if len(t) > len(subs) {
				subs = t
			}
		}
	}

	if even && len(s[i-epvt+1:i]) > len(subs) {
		subs = s[i-epvt+1 : i]
	}

	if odd && len(s[i-opvt+1:i]) > len(subs) {
		subs = s[i-opvt+1 : i]
	}

	return subs
}

func findPalindrome(s string) string {
	l := len(s)

	if l <= 1 {
		return s
	}

	if l == 2 {
		if s[0] != s[1] {
			return string(s[0])
		} else {
			return s
		}
	}

	left, right := -1, -1
	i, k := 0, l-1
	for i = 0; i < l && i <= k; i++ {
		if s[i] == s[k] {
			if left == -1 || right == -1 {
				left, right = i, k
			}
		} else {
			left, right = -1, -1
		}
		k--
	}

	if left != -1 && right != -1 {
		return s[left : right+1]
	}

	return ""
}
