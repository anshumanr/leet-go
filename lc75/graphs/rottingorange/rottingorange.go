package rottingorange

type pos struct {
	x, y, steps int
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	steps := 0

	rotten := []pos{}
	for x, a := range grid {
		for y, b := range a {
			if b == 2 {
				rotten = append(rotten, pos{x, y, 0})
			}
		}
	}

	for len(rotten) > 0 {
		curr := rotten[0]
		rotten = rotten[1:]
		x, y := curr.x, curr.y

		//above
		if x-1 >= 0 && grid[x-1][y] == 1 {
			steps = curr.steps + 1
			rotten = append(rotten, pos{x - 1, y, curr.steps + 1})
			grid[x-1][y] = 2
		}
		//below
		if x+1 < m && grid[x+1][y] == 1 {
			steps = curr.steps + 1
			rotten = append(rotten, pos{x + 1, y, curr.steps + 1})
			grid[x+1][y] = 2
		}
		//right
		if y+1 < n && grid[x][y+1] == 1 {
			steps = curr.steps + 1
			rotten = append(rotten, pos{x, y + 1, curr.steps + 1})
			grid[x][y+1] = 2
		}
		//left
		if y-1 >= 0 && grid[x][y-1] == 1 {
			steps = curr.steps + 1
			rotten = append(rotten, pos{x, y - 1, curr.steps + 1})
			grid[x][y-1] = 2
		}
	}

	for _, a := range grid {
		for _, b := range a {
			if b == 1 {
				return -1
			}
		}
	}

	return steps
}

func orangesRotting_0ms(grid [][]int) int {
	var d [][2]int
	moves := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				d = append(d, [2]int{i, j})
			}
		}
	}
	t := 0
	for len(d) > 0 {
		N := len(d)
		for i := 0; i < N; i++ {
			x := d[0][0]
			y := d[0][1]
			d = d[1:]
			for j := 0; j < 4; j++ {
				u := x + moves[j][0]
				v := y + moves[j][1]
				// fmt.Println(u,v,m,n)
				if u >= 0 && v >= 0 && u < m && v < n {
					// fmt.Println(u,v)
					if grid[u][v] == 1 {
						grid[u][v] = 2
						d = append(d, [2]int{u, v})
					}
				}
			}
			// fmt.Println(d)
		}
		t++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if t > 0 {
		return t - 1
	}
	return t
}
