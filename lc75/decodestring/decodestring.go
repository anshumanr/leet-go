package decodestring

func decodeString(s string) string {
	l := len(s)

	mult, i, j, ret := 0, 0, 0, []byte{}
	for i < l {
		switch {
		case s[i] >= 48 && s[i] <= 57:
			mult = 10*mult + int(s[i]-48)
			i++
			continue

		case mult != 0:
			match := 0
			for j = i; j < l; j++ {
				if s[j] == '[' {
					match++
				}
				if s[j] == ']' {
					match--
					if match == 0 {
						break
					}
				}
			}
			dec := decodeString(s[i+1 : j])
			for k := 0; k < mult; k++ {
				ret = append(ret, []byte(dec)...)
			}
			i = j + 1
			mult = 0
			continue

		default:
			ret = append(ret, s[i])
			i++
		}
	}

	return string(ret)
}

func decodeString_leet0ms(s string) string {
	var stack []*struct {
		d int
		s []uint8
	}

	var result []uint8

	isDigit := func(c uint8) bool { return c >= '0' && c <= '9' }

	var si *struct {
		d int
		s []uint8
	}
	var digit int

	for i := range s {

		if isDigit(s[i]) {
			digit = digit*10 + int(s[i]-'0')
			continue
		}

		if s[i] == '[' {
			if si != nil {
				stack = append(stack, si)
			}
			si = &struct {
				d int
				s []uint8
			}{d: digit, s: []uint8{}}
			digit = 0
			continue
		}

		if s[i] == ']' {
			if len(stack) == 0 {
				for si.d > 0 {
					result = append(result, si.s...)
					si.d--
				}
				si = nil
			} else {
				for si.d > 0 {
					stack[len(stack)-1].s = append(stack[len(stack)-1].s, si.s...)
					si.d--
				}
				si = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			continue
		}

		if si != nil {
			(*si).s = append((*si).s, s[i])
			continue
		}

		result = append(result, s[i])
	}

	return string(result)
}
