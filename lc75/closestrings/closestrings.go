package closestrings

import "sort"

func closeStrings(word1 string, word2 string) bool {
	l1, l2 := len(word1), len(word2)
	if l1 != l2 {
		return false
	}
	arr1, arr2 := make([]int, 26), make([]int, 26)
	for i := 0; i < l1; i++ {
		arr1[word1[i]-'a']++
		arr2[word2[i]-'a']++
	}

	arl1 := len(arr1)
	for j := 0; j < arl1; j++ {
		if arr1[j] == 0 && arr2[j] == 0 {
			continue
		}
		if (arr1[j] != 0 && arr2[j] == 0) || (arr2[j] != 0 && arr1[j] == 0) {
			return false
		}
	}

	sort.Ints(arr1)
	sort.Ints(arr2)
	for j := 0; j < arl1; j++ {
		if arr1[j] == 0 && arr2[j] == 0 {
			continue
		}
		if arr1[j] != arr2[j] {
			return false
		}
	}

	return true
}

func charCount(s string) []int {
	arr := make([]int, 26)
	for i := 0; i < len(s); i = i + 1 {
		arr[int(s[i]-'a')]++
	}
	return arr
}

func closeStrings_leet16ms(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := charCount(word1)
	arr2 := charCount(word2)
	for i := 0; i < len(arr1); i++ {
		if (arr1[i] > 0 && arr2[i] == 0) || (arr2[i] > 0 && arr1[i] == 0) {
			return false
		}
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
