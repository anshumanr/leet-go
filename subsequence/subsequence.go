package subsequence

func isSubsequence(s string, t string) bool {
	sl := len(s)
	tl := len(t)

	if sl == 0 {
		return true
	}

	found := false
	k := 0
	for i := 0; i < sl; i++ {
		found = false
		for ; k < tl; k++ {
			if s[i] == t[k] {
				k++
				found = true
				break
			}
		}
		if k == tl && !found {
			return false
		}
	}

	return found
}

func isSubsequence_leet0ms(s string, t string) bool {
	s_index := 0
	t_index := 0

	for s_index < len(s) {
		if (t_index) == len(t) {
			break
		}

		if s[s_index] == t[t_index] {
			s_index++
		}

		t_index++
	}

	return s_index == len(s)
}
