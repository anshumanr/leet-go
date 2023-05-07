package compress

import (
	"strconv"
)

func compress(chars []byte) int {
	n := 0
	l := len(chars)
	if l <= 1 {
		return l
	}

	i, cnt, prev := 0, 0, chars[0]
	for ; i < l; i++ {
		if chars[i] == prev {
			cnt++
		} else {
			chars[n] = prev
			n++
			switch {
			case cnt > 1 && cnt < 10:
				chars[n] = byte(cnt + 48)
				n++
			case cnt >= 10:
				t := strconv.Itoa(cnt)
				for _, v := range t {
					chars[n] = byte(v)
					n++
				}
			}
			prev, cnt = chars[i], 1
		}
	}

	chars[n] = prev
	n++
	switch {
	case cnt > 1 && cnt < 10:
		chars[n] = byte(cnt + 48)
		n++
	case cnt >= 10:
		t := strconv.Itoa(cnt)
		for _, v := range t {
			chars[n] = byte(v)
			n++
		}
	}

	return n
}

func compress_leet0ms(chars []byte) int {
	wi, chCnt, lCh := 0, 1, chars[0]
	for ri, ch := range chars {
		if ri == 0 {
			continue
		}

		if ch == lCh {
			chCnt++
			continue
		}

		wi = writeCharsAndMoveIndex(lCh, chCnt, wi, &chars)
		chCnt = 1
		lCh = ch
	}
	wi = writeCharsAndMoveIndex(lCh, chCnt, wi, &chars)

	chars = chars[:wi]
	return len(chars)
}

func writeCharsAndMoveIndex(ch byte, chCnt, writeIdx int, chars *[]byte) int {
	// write ch to chars
	(*chars)[writeIdx] = ch
	writeIdx++

	// write chCnt to chars
	if chCnt > 1 {
		cntBytes := []byte(strconv.Itoa(chCnt))
		for _, b := range cntBytes {
			(*chars)[writeIdx] = b
			writeIdx++
		}
	}

	// return updated write index
	return writeIdx
}
