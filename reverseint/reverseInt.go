package reverseint

import "strconv"

const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

func reverse(x int) int {
	if x > maxInt32 || x < minInt32 {
		return 0
	}

	revInt := 1
	sign := 1
	if x < 0 {
		sign = -1
		x = x * sign
	}

	r := x % 10
	revInt *= r

	m := 10
	x = x / 10
	for ; x != 0; x = x / 10 {
		r := x % 10
		revInt = revInt*m + r

		//check
		if revInt > maxInt32 {
			return 0
		}

		if revInt < minInt32 {
			return 0
		}
	}

	return revInt * sign

}

func reverse_leet0ms(x int) int {
	numStr := strconv.Itoa(x)
	reversedStr := ""
	signed := false
	for _, digitStr := range numStr {
		if digitStr == '-' {
			signed = true
		} else {
			reversedStr = string(digitStr) + reversedStr
		}
	}
	if signed {
		reversedStr = "-" + reversedStr
	}
	reversedInt, err := strconv.ParseInt(reversedStr, 10, 32)
	if err != nil {
		return 0
	}
	return int(reversedInt)
}
