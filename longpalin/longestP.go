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
	subs := string(s[0])

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
	case 3:
		if s[0] == s[2] {
			return s
		} else if s[0] == s[1] {
			return s[:2]
		} else if s[1] == s[2] {
			return s[1:]
		} else {
			return string(s[0])
		}
	}

	epvt := 0
	opvt := 1
	even, odd := false, false
	oleft, oright := -1, -1
	eleft, eright := -1, -1
	i := 2
	for i = 1; i < l; i++ {
		//even length palindrome
		for i-epvt >= 0 && i+epvt+1 < l {
			if s[i-epvt] == s[i+1+epvt] {
				eleft = i - epvt
				eright = i + 1 + epvt
				even = true
			} else if even {
				t := s[eleft : eright+1]
				if len(t) > len(subs) {
					subs = t
				}
				epvt = 0
				even = false
				eleft, eright = -1, -1
				break
			} else {
				break
			}
			epvt++
		}

		if even && eleft != -1 && eright != -1 {
			if eright-eleft == l {
				return s[eleft : eright+1]
			}

			t := s[eleft : eright+1]
			if len(t) > len(subs) {
				subs = t
			}
			even = false
			epvt = 0
			eleft, eright = -1, -1
		}

		//odd length palindrome
		for i-opvt >= 0 && i+opvt < l {
			if s[i+opvt] == s[i-opvt] {
				oleft = i - opvt
				oright = i + opvt
				odd = true
			} else if odd {
				t := s[oleft : oright+1]
				if len(t) > len(subs) {
					subs = t
				}
				odd = false
				opvt = 1
				oleft, oright = -1, -1
				break
			} else {
				break
			}
			opvt++
		}

		if odd && oleft != -1 && oright != -1 {
			if oright-oleft == l {
				return s[eleft : eright+1]
			}

			t := s[oleft : oright+1]
			if len(t) > len(subs) {
				subs = t
			}
			odd = false
			opvt = 1
			oleft, oright = -1, -1
		}
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

func longestPalindrome0ms(s string) string {
	l := len(s)
	var maxStart, maxLen int

	for i := 0; i < l; {
		left, right := i, i
		for right < l-1 && s[right] == s[right+1] {
			right++
		}

		i = right + 1

		for left > 0 && right < l-1 && s[left-1] == s[right+1] {
			left--
			right++
		}

		if maxLen < (right - left) {
			maxStart = left
			maxLen = right - left
		}
	}

	return s[maxStart : maxLen+maxStart+1]
}

func longestPalindrome4ms(s string) string {
	pp := ""

	for i := 0; i < len(s); i++ {
		p1 := getPalindrome(s, i, i)
		p2 := getPalindrome(s, i, i+1)

		x := p2
		if len(p1) > len(p2) {
			x = p1
		}

		if len(x) > len(pp) {
			pp = x
		}

	}

	return pp

}

func getPalindrome(s string, l int, r int) string {

	ss := ""
	for i, j := l, r; j < len(s) && i >= 0; i, j = i-1, j+1 {
		if s[i] != s[j] {
			return ss
		} else {
			ss = s[i : j+1]
		}

	}

	return ss
}
