package floodfill

func floodFill(image [][]int, row, col, color int) [][]int {
	if image[row][col] == color {
		return image
	}

	m := len(image)
	n := len(image[0])

	visited := make([]bool, m*n)
	stack := make([]int, m*n)
	top := 0

	stack[top] = row*n + col
	top++

	for top > 0 {
		top--
		pix := stack[top]

		c := pix % n
		r := pix / n
		old := image[r][c]
		image[r][c] = color
		visited[pix] = true

		//top
		if r-1 >= 0 && image[r-1][c] == old && !visited[(r-1)*n+c] {
			stack[top] = (r-1)*n + c
			top++
		}

		//below
		if r+1 < m && image[r+1][c] == old && !visited[(r+1)*n+c] {
			stack[top] = (r+1)*n + c
			top++
		}

		//left
		if c-1 >= 0 && image[r][c-1] == old && !visited[r*n+c-1] {
			stack[top] = r*n + c - 1
			top++
		}

		//right
		if c+1 < n && image[r][c+1] == old && !visited[r*n+c+1] {
			stack[top] = r*n + c + 1
			top++
		}
	}

	return image
}

func floodFillv2(image [][]int, row, col, color int) [][]int {
	if image[row][col] == color {
		return image
	}
	return dfsv2(image, row, col, color, image[row][col])
}

func dfsv2(image [][]int, row, col, color, old int) [][]int {
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) {
		return image
	}

	if image[row][col] == color || image[row][col] != old {
		return image
	}

	image[row][col] = color

	//top
	dfsv2(image, row-1, col, color, old)
	//bottom
	dfsv2(image, row+1, col, color, old)
	//left
	dfsv2(image, row, col-1, color, old)
	//right
	dfsv2(image, row, col+1, color, old)

	return image
}
