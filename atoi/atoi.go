package atoi

import (
	"math"
	"strings"
	"unicode"
)

var atoiMap = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

const (
	maxVal = 2147483647
	minVal = -2147483648
)

func myAtoi1(str string) int {
	s := strings.Trim(str, " ")

	if ok := sanityCheck(s); !ok {
		return 0
	}

	intVal := 0
	minus := 1

	for i := range s {
		if i == 0 {
			if s[i] == '-' {
				minus = -1
				continue
			}

			if s[i] == '+' {
				continue
			}
		}

		//fmt.Println(intVal, s[i])
		if s[i] > 47 && s[i] < 58 {
			intVal = intVal*10 + atoiMap[s[i]]

			if intVal > maxVal || intVal < minVal {
				break
			}
		} else {
			break
		}
	}

	if intVal*minus > maxVal {
		return maxVal
	} else if intVal*minus < minVal {
		return minVal
	}

	return intVal * minus
}

func myAtoi2(str string) int {

	intVal := 0
	minus := 1

	flag := false
	for i := range str {
		if flag == false {
			if str[i] == ' ' {
				continue
			}

			flag = true

			if ok := sanityCheck(str[i:]); !ok {
				return 0
			}

			if str[i] == '-' {
				minus = -1
				continue
			}

			if str[i] == '+' {
				continue
			}
		}

		if str[i] > 47 && str[i] < 58 {
			intVal = intVal*10 + atoiMap[str[i]]
			if intVal > maxVal || intVal < minVal {
				break
			}
		} else {
			break
		}
	}

	if intVal*minus > maxVal {
		return maxVal
	} else if intVal*minus < minVal {
		return minVal
	}

	return intVal * minus
}

func myAtoi3(str string) int {

	intVal := 0
	minus := 1

	flag := false
	for i, c := range str {
		if flag == false {
			if c == ' ' {
				continue
			}

			flag = true

			if ok := sanityCheck(str[i:]); !ok {
				return 0
			}

			if c == '-' {
				minus = -1
				continue
			}

			if c == '+' {
				continue
			}
		}

		if str[i] > 47 && str[i] < 58 {
			intVal = intVal*10 + int(c-'0')
			if intVal > maxVal || intVal < minVal {
				break
			}
		} else {
			break
		}
	}

	if intVal*minus > maxVal {
		return maxVal
	} else if intVal*minus < minVal {
		return minVal
	}

	return intVal * minus
}

func sanityCheck(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		if len(s) < 2 {
			return false
		}

		if s[1] > 47 && s[1] < 58 {
			return true
		}
	} else if s[0] > 47 && s[0] < 58 {
		return true
	}

	return false
}

func atoiLeet0ms(str string) int {
	// discard whitespace
	start := 0
	for i, c := range str {
		if !unicode.IsSpace(c) {
			start = i
			break
		}
	}
	if str == "" {
		return start
	}
	// is negative
	neg := str[start] == '-'
	if str[start] == '-' || str[start] == '+' {
		start = start + 1
	}

	res := 0
	for _, c := range str[start:] {
		if !unicode.IsDigit(c) {
			break
		}
		res = res * 10
		res = res + int(c-'0')
		if neg && res > math.MaxInt32+1 {
			return math.MinInt32
		} else if !neg && res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if neg {
		return -1 * res
	}
	return res
}
