package exitmaze

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	paths := [][]int{}
	visited := make([]bool, m*n)

	si, sj := entrance[0], entrance[1]
	paths = append(paths, []int{si, sj})
	visited[si*n+sj] = true
	steps := 0

	for {
		l := len(paths)
		if l == 0 {
			break
		}

		inc := false
		for ind := 0; ind < l; ind++ {
			i, j := paths[ind][0], paths[ind][1]

			if (i == 0 || i == m-1 || j == 0 || j == n-1) && (i != si || j != sj) {
				return steps
			}

			//above
			if i-1 >= 0 && maze[i-1][j] == '.' && !visited[(i-1)*n+j] {
				paths = append(paths, []int{i - 1, j})
				visited[(i-1)*n+j] = true
				inc = true
			}
			//below
			if i+1 < m && maze[i+1][j] == '.' && !visited[(i+1)*n+j] {
				paths = append(paths, []int{i + 1, j})
				visited[(i+1)*n+j] = true
				inc = true
			}
			//right
			if j+1 < n && maze[i][j+1] == '.' && !visited[i*n+j+1] {
				paths = append(paths, []int{i, j + 1})
				visited[i*n+j+1] = true
				inc = true
			}
			//left
			if j-1 >= 0 && maze[i][j-1] == '.' && !visited[i*n+j-1] {
				paths = append(paths, []int{i, j - 1})
				visited[i*n+j-1] = true
				inc = true
			}
		}
		if inc {
			steps++
		}
		paths = paths[l:]
	}

	return -1
}

type pos struct {
	i, j int
}

func nearestExitV2(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	paths := []pos{}

	si, sj := entrance[0], entrance[1]
	paths = append(paths, pos{si, sj})
	steps := 0

	for {
		l := len(paths)
		if l == 0 {
			break
		}

		inc := false
		for ind := 0; ind < l; ind++ {
			cur := paths[ind]
			i, j := cur.i, cur.j

			if (i == 0 || i == m-1 || j == 0 || j == n-1) && (i != si || j != sj) {
				return steps
			}

			//above
			if i-1 >= 0 && maze[i-1][j] == '.' {
				paths = append(paths, pos{i - 1, j})
				maze[i-1][j] = '+'
				inc = true
			}
			//below
			if i+1 < m && maze[i+1][j] == '.' {
				paths = append(paths, pos{i + 1, j})
				maze[i+1][j] = '+'
				inc = true
			}
			//right
			if j+1 < n && maze[i][j+1] == '.' {
				paths = append(paths, pos{i, j + 1})
				maze[i][j+1] = '+'
				inc = true
			}
			//left
			if j-1 >= 0 && maze[i][j-1] == '.' {
				paths = append(paths, pos{i, j - 1})
				maze[i][j-1] = '+'
				inc = true
			}
		}
		if inc {
			steps++
		}
		paths = paths[l:]
	}

	return -1
}

type Position struct {
	x, y, steps int
}

func isValid(maze [][]byte, x, y int) bool {
	if x < 0 || x >= len(maze) {
		return false
	}
	if y < 0 || y >= len(maze[0]) {
		return false
	}
	if maze[x][y] != '.' {
		return false
	}
	return true
}

func isExit(x, y int, entrance []int, rows, cols int) bool {
	if x == entrance[0] && y == entrance[1] {
		return false
	}
	if x > 0 && x < rows-1 && y > 0 && y < cols-1 {
		return false
	}
	return true
}

func nearestExit_leet6ms(maze [][]byte, e []int) int {
	rows, cols := len(maze), len(maze[0])
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Position{{e[0], e[1], 0}}
	// visited := make(map[[2]int]bool)
	// visited[[2]int{e[0], e[1]}] = true
	maze[e[0]][e[1]] = '+'
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			nx, ny := cur.x+d[0], cur.y+d[1]
			if isValid(maze, nx, ny) {
				if isExit(nx, ny, e, rows, cols) {
					return cur.steps + 1
				}
				queue = append(queue, Position{nx, ny, cur.steps + 1})
				maze[nx][ny] = '+'
			}
		}
	}
	return -1
}
