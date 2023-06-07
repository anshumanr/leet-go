package reorderroutes

func minReorder(n int, connections [][]int) int {
	change := 0
	adjlist := make([][]int, n)
	visited := make([]bool, n)

	for _, ab := range connections {
		adjlist[ab[0]] = append(adjlist[ab[0]], ab[1])
		adjlist[ab[1]] = append(adjlist[ab[1]], -ab[0])
	}

	dfs(0, adjlist, visited, &change)

	return change
}

func dfs(start int, adjlist [][]int, visited []bool, change *int) {
	visited[start] = true
	nbr := adjlist[start]
	for _, v := range nbr {
		abv := abs(v)
		if !visited[abv] {
			if v == abv {
				*change++
			}
			dfs(abv, adjlist, visited, change)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minReorder_leet149ms(n int, connections [][]int) int {
	// Create an adjacency list to represent the tree
	adj := make([][]int, n)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], ^a) // Use negative sign to indicate a reverse edge
	}

	// Perform a depth-first search from city 0, counting the number of reverse edges
	count := 0
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if v >= 0 && !visited[v] {
				count++
				dfs(v)
			} else if v < 0 && !visited[^v] {
				dfs(^v)
			}
		}
	}
	dfs(0)

	return count
}
