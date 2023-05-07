package reversevowels

func reverseVowels(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}

	b := []byte(s)

	i, j := 0, l-1
	for i < j {
		if isVowel(s[i]) && isVowel(s[j]) {
			b[i], b[j] = s[j], s[i]
			j--
			i++
		}
		if !isVowel(s[j]) {
			j--
			continue
		}
		if !isVowel(s[i]) {
			i++
		}
	}

	return string(b)
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U'
}

func reverseVowels_leet0ms(s string) string {
	arr := []rune(s)

	v := []int{}
	for i := range s {
		if isVowel2(arr[i]) {
			v = append(v, i)
		}
	}

	i, j := 0, len(v)-1
	for i < j {
		i1 := v[i]
		i2 := v[j]
		arr[i1], arr[i2] = arr[i2], arr[i1]
		i++
		j--
	}

	return string(arr)
}

func isVowel2(r rune) bool {
	switch r {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	case 'A':
		return true
	case 'E':
		return true
	case 'I':
		return true
	case 'O':
		return true
	case 'U':
		return true
	default:
		return false
	}
	return false
}
