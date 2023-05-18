package dota2

func predictPartyVictory(senate string) string {
	sq := []byte(senate)

	i, ri, di := 0, 0, 0
	for len(sq) > 0 {
		if i >= len(sq) {
			i, ri, di = 0, 0, 0
		}
		if sq[i] == 'R' {
			di = find(sq, 'D', i)
			if di >= 0 {
				sq = append(sq[:di], sq[di+1:]...)
				i++
			} else {
				return "Radiant"
			}
			continue
		}
		if sq[i] == 'D' {
			ri = find(sq, 'R', i)
			if ri >= 0 {
				sq = append(sq[:ri], sq[ri+1:]...)
				i++
			} else {
				return "Dire"
			}
			continue
		}
	}

	return string(sq)
}

func find(arr []byte, ch byte, k int) int {
	l := len(arr)
	for i := k; i < l; i++ {
		if arr[i] == ch {
			return i
		}
	}
	for i := 0; i < k; i++ {
		if arr[i] == ch {
			return i
		}
	}
	return -1
}

func predictPartyVictory_leet4ms(senate string) string {
	R := make([]int, 0, len(senate))
	D := make([]int, 0, len(senate))

	for i, r := range senate {
		if r == 'R' {
			R = append(R, i)
		} else {
			D = append(D, i)
		}
	}

	for len(D) > 0 && len(R) > 0 {
		ld := D[0]
		D = D[1:]

		lr := R[0]
		R = R[1:]

		if lr < ld {
			R = append(R, lr+len(senate))
		} else {
			D = append(D, ld+len(senate))
		}
	}

	if len(R) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}
