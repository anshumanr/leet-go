package longestpalindrome

/*
  Given a string s which consists of lowercase or uppercase letters,
  return the length of the longest palindrome that can be built with those letters.
*/
func longestPalindrome(s string) int {
	var arr [128]int
	odd := (len(s)%2 == 1)
	oddOccur, longest := 0, len(s)

	for _, v := range s {
		arr[v]++
		if arr[v]%2 == 0 {
			oddOccur--
		} else {
			oddOccur++
		}
	}

	if odd || oddOccur != 0 {
		return longest - oddOccur + 1
	}

	return longest
}

func longestPalindrome_leet0ms(s string) int {
	m := make(map[byte]int)
	var oddCount int
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for _, val := range m {
		if val%2 != 0 {
			oddCount++
		}
	}
	if oddCount > 0 {
		return len(s) - oddCount + 1
	}
	return len(s) - oddCount
}
