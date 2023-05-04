package reversewords

import "strings"

/*
  You are given two strings word1 and word2.
  Merge the strings by adding letters in alternating order, starting with word1.
  If a string is longer than the other, append the additional letters onto the
  end of the merged string. Return the merged string.
*/
func reverseWords(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}

	s = strings.Trim(s, " ")
	l = len(s)

	var a string
	i := l - 1
	wb, space := true, false
	for ; i > 0; i-- {
		if s[i] == ' ' {
			if wb {
				continue
			} else {
				if space {
					a = a + " " + s[i+1:]
				} else {
					a = a + s[i+1:]
				}
				s = strings.Trim(s[:i], " ")
				i = len(s)
				wb, space = true, true
			}
		} else {
			wb = false
		}
	}

	if !wb || s[0] != ' ' {
		if space {
			a = a + " " + s[i:]
		} else {
			a = a + s[i:]
		}
	}

	return a
}

func reverseWordsV2(s string) string {
	l := len(s)
	if l == 1 {
		return s
	}

	s = strings.TrimSpace(s)
	l = len(s)

	var a strings.Builder
	a.Grow(l)
	i := l - 1
	wb, space := true, false
	for ; i > 0; i-- {
		if s[i] == ' ' {
			if wb {
				continue
			} else {
				if space {
					a.WriteString(" " + s[i+1:])
				} else {
					a.WriteString(s[i+1:])
				}
				s = strings.TrimSpace(s[:i])
				i = len(s)
				wb, space = true, true
			}
		} else {
			wb = false
		}
	}

	if !wb || s[0] != ' ' {
		if space {
			a.WriteString(" " + s[i:])
		} else {
			a.WriteString(s[i:])
		}
	}

	return a.String()
}

func reverseWords_leet0ms(s string) string {
	words := strings.Fields(s)

	reverseSlice := func(s []string) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	reverseSlice(words)

	return strings.Join(words, " ")
}
