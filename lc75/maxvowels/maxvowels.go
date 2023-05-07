package maxvowels

import "strings"

func maxVowels(s string, k int) int {
	if strings.IndexAny(s, "aeiou") < 0 {
		return 0
	}

	i, l, cnt := 0, len(s), 0
	for ; i < l-k+1; i++ {
		t := 0
		if c := strings.Count(s[i:i+k], "a"); c > 0 {
			t += c
		}
		if c := strings.Count(s[i:i+k], "e"); c > 0 {
			t += c
		}
		if c := strings.Count(s[i:i+k], "i"); c > 0 {
			t += c
		}
		if c := strings.Count(s[i:i+k], "o"); c > 0 {
			t += c
		}
		if c := strings.Count(s[i:i+k], "u"); c > 0 {
			t += c
		}

		if t >= k {
			return k
		} else if t > cnt {
			cnt = t
		}
	}
	return cnt
}

func maxVowelsV3(s string, k int) int {
	if strings.IndexAny(s, "aeiou") < 0 {
		return 0
	}

	t := count(s[:k])

	i, l, max := k, len(s), t
	for ; i < l; i++ {
		if isVowel(s[i-k]) {
			t--
		}
		if isVowel(s[i]) {
			t++
		}
		if t >= k {
			return k
		}
		if t > max {
			max = t
		}
	}
	return max
}

func count(s string) int {
	cnt := 0
	for _, v := range s {
		if isVowel(byte(v)) {
			cnt++
		}
	}
	return cnt
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}

func maxVowels_leet0ms(s string, k int) int {
	vowels := 0
	i, j := 0, 0
	max := -1

	for j < len(s) {
		if isVowel2(s[j]) {
			vowels++
		}

		if j-i+1 == k {
			if vowels > max {
				max = vowels
			}

			if isVowel2(s[i]) {
				vowels--
			}

			i++
		}

		j++
	}

	return max
}

func isVowel2(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
