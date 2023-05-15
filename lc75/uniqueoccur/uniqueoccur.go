package increasingtriplet

func uniqueOccurrences(arr []int) bool {
	l := len(arr)
	cm := make(map[int]int)

	for i := 0; i < l; i++ {
		cm[arr[i]]++
	}

	uniq := make(map[int]bool, len(cm))
	for _, v := range cm {
		if uniq[v] {
			return false
		}
		uniq[v] = true
	}

	return true
}

func uniqueOccurrences_leet0ms(arr []int) bool {
	occurrenceMap := map[int]int{}
	for _, i := range arr {
		if _, ok := occurrenceMap[i]; !ok {
			occurrenceMap[i] = 0
		}
		occurrenceMap[i]++
	}

	occurrenceOccurMap := map[int]struct{}{}
	for _, occurrence := range occurrenceMap {
		if _, ok := occurrenceOccurMap[occurrence]; ok {
			return false
		}
		occurrenceOccurMap[occurrence] = struct{}{}
	}

	return true
}
