package removestars

func removeStars(s string) string {
	l := len(s)
	byt := []byte(s)

	for i := 0; i < l; {
		if byt[i] == '*' {
			if i < l-1 {
				byt = append(byt[:i-1], byt[i+1:]...)
				i = i - 1
				l -= 2
			} else {
				byt = byt[:l-2]
				break
			}
		} else {
			i++
		}
	}

	return string(byt)
}

func removeStarsV2(s string) string {
	l := len(s)
	byt := make([]byte, l)

	i, j := 0, 0
	for i < l {
		if s[i] == '*' {
			j--
		} else {
			byt[j] = byte(s[i])
			j++
		}
		i++
	}

	return string(byt[:j])
}

func removeStars_leet32ms(s string) string {
	var resp []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			resp = resp[0 : len(resp)-1]
		} else {
			resp = append(resp, s[i])
		}
	}

	return string(resp)

}
