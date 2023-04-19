package isomorphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n1 := mark(s)
	return n1 == mark(t)
}

func mark(s string) int {
	m := make(map[byte]int, len(s))

	num := 0
	digit := 1
	for _, v := range s {
		n, ok := m[byte(v)]
		if !ok {
			m[byte(v)] = digit
			num = num*10 + digit
			digit++
		} else {
			num = num*10 + n
		}
	}

	return num
}

func isIsomorphicV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var m1 [256]int
	var m2 [256]int

	num1 := 0
	num2 := 0
	digit1 := 1
	digit2 := 1
	for i, _ := range s {
		if m1[int(s[i])] == 0 {
			m1[int(s[i])] = digit1
			num1 = num1*10 + digit1
			digit1++
		} else {
			num1 = num1*10 + m1[int(s[i])]
		}

		if m2[int(t[i])] == 0 {
			m2[int(t[i])] = digit2
			num2 = num2*10 + digit2
			digit2++
		} else {
			num2 = num2*10 + m2[int(t[i])]
		}
	}

	return num1 == num2
}

func isIsomorphic_leet1ms(s string, t string) bool {
	var sarr [300]int
	var tarr [300]int

	// for i := 0; i < len(sarr); i++ {
	// 	sarr[i] = 0
	// }
	// for i := 0; i < len(tarr); i++ {
	// 	tarr[i] = 0
	// }

	for i := 0; i < len(s); i++ {
		if sarr[s[i]] != 0 {
			if t[i] != t[sarr[s[i]]-1] {
				return false
			}
		}
		sarr[s[i]] = i + 1

		if tarr[t[i]] != 0 {
			if s[i] != s[tarr[t[i]]-1] {
				return false
			}
		}
		tarr[t[i]] = i + 1
	}

	return true
}
