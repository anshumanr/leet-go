package mergealternately

/*
  You are given two strings word1 and word2.
  Merge the strings by adding letters in alternating order, starting with word1.
  If a string is longer than the other, append the additional letters onto the
  end of the merged string. Return the merged string.
*/
func mergeAlternately(word1 string, word2 string) string {
	lw, sw := word1, word2

	if len(word1) < len(word2) {
		sw, lw = word1, word2
	}

	l1, l2 := len(lw), len(sw)

	ans := make([]byte, l1+l2)

	i, j := 0, 0
	for ; i < l2; i++ {
		ans[j] = word1[i]
		ans[j+1] = word2[i]
		j += 2
	}

	for i < l1 {
		ans[j] = lw[i]
		i++
		j++
	}

	return string(ans)
}

func mergeAlternately_leet0ms(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	result := make([]byte, l1+l2)

	for i := 0; i < l1 && i < l2; i++ {
		result[i*2] = word1[i]
		result[i*2+1] = word2[i]
	}
	for i := l1; i < l2; i++ {
		result[l1+i] = word2[i]
	}
	for i := l2; i < l1; i++ {
		result[l2+i] = word1[i]
	}

	return string(result)
}
