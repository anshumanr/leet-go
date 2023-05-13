package equalpairs

func equalPairs(grid [][]int) int {
	l := len(grid)
	//rowarr := make([]int, l)
	colarr := make([][]int, l)

	for i := 0; i < l; i++ {
		col := make([]int, l)
		for k := 0; k < l; k++ {
			col[k] = grid[k][i]
		}
		colarr[i] = col
	}

	max := 0
	for _, v := range grid {
		for _, w := range colarr {
			if compare(v, w) {
				max++
			}
		}
	}

	return max
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalPairs_leet71ms(grid [][]int) int {
	rows := make(map[int]int, len(grid))
	for _, r := range grid {
		v := 0
		for _, c := range r {
			v = v*10 + c
		}
		rows[v]++
	}

	res := 0
	for i := 0; i < len(grid[0]); i++ {
		v := 0
		for j := 0; j < len(grid); j++ {
			v = v*10 + grid[j][i]
		}
		res += rows[v]
	}
	return res
}
