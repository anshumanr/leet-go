package parentheses

func longestValidParentheses(s string) int {
	longest := 0
	slen := len(s)
	if slen < 2 {
		return longest
	}

	sbyte := []byte(s)
	stack := []int{}
	for i := 0; i < slen; i++ {
		if s[i] == '(' {
			//push
			stack = append(stack, i)
		} else {
			//pop
			l := len(stack)
			if l > 0 && s[stack[l-1]] == '(' {
				sbyte[i] = '1'
				sbyte[stack[l-1]] = '1'
				stack = stack[:l-1]
			}
		}
	}

	tl := 0
	for _, v := range sbyte {
		if v == '1' {
			tl++
		} else {
			if tl > longest {
				longest = tl
			}
			tl = 0
		}
	}

	if tl > longest {
		longest = tl
	}

	return longest
}

func longestValidParentheses2(s string) int {
	longest := 0
	slen := len(s)
	if slen < 2 {
		return longest
	}

	sbyte := []byte(s)
	stack := make([]int, slen)
	sp := 0
	for i := 0; i < slen; i++ {
		if s[i] == '(' {
			//push
			stack[sp] = i
			sp++
		} else {
			//pop
			if sp > 0 && s[stack[sp-1]] == '(' {
				sbyte[i] = '1'
				sbyte[stack[sp-1]] = '1'
				sp--
			}
		}
	}

	tl := 0
	for _, v := range sbyte {
		if v == '1' {
			tl++
		} else {
			if tl > longest {
				longest = tl
			}
			tl = 0
		}
	}

	if tl > longest {
		longest = tl
	}

	return longest
}

func longestValidParentheses_leet(s string) int {
	if len(s) == 0 {
		return 0
	}
	n, left, max := len(s), 0, 0
	d := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else if left > 0 {
			d[i] = d[i-1] + 2
			if i-d[i] > 0 {
				d[i] += d[i-d[i]]
			}
			max = Max(max, d[i])
			left--
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
