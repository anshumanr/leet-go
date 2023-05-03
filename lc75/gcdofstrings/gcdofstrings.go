package gcdofstrings

/*
  You are given two strings word1 and word2.
  Merge the strings by adding letters in alternating order, starting with word1.
  If a string is longer than the other, append the additional letters onto the
  end of the merged string. Return the merged string.
*/
func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 == l2 && str1 != str2 {
		return ""
	}

	g := gcd(l1, l2)

	if str1[:g] != str2[:g] {
		return ""
	}

	i := g
	for ; i < l1; i += g {
		if str1[:g] != str1[i:i+g] {
			return ""
		}
	}

	i = g
	for ; i < l2; i += g {
		if str2[:g] != str2[i:i+g] {
			return ""
		}
	}

	return str1[:g]

}

func gcdOfStringsV2(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 == l2 && str1 != str2 {
		return ""
	}
	if l1 > l2 {
		str1, str2 = str2, str1
		l1, l2 = l2, l1
	}

	g := gcd(l1, l2)

	if str1[:g] != str2[:g] {
		return ""
	}

	i := g
	for ; i < l1; i += g {
		if str1[i:i+g] != str2[i:i+g] {
			return ""
		}
	}

	for ; i < l2; i += g {
		if str2[:g] != str2[i:i+g] {
			return ""
		}
	}

	return str1[:g]

}

//function to find GCD of two numbers
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func gcdOfStrings_leet0ms(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		return gcdOfStrings_leet0ms(str2, str1)
	}

	if !starts(str1, str2) {
		return ""
	}

	if len(str2) == 0 {
		return str1
	}

	return gcdOfStrings_leet0ms(str1[len(str2):], str2)
}

func starts(str, prefix string) bool {
	for i := range prefix {
		if str[i] != prefix[i] {
			return false
		}
	}
	return true
}
