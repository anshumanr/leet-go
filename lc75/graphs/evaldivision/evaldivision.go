package evaldivision

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n := len(equations)
	graph := make(map[string][]string)
	edgelen := make(map[string]float64, n)

	for ind, v := range equations {
		if _, ok := edgelen[v[0]+"-"+v[1]]; !ok {
			graph[v[0]] = append(graph[v[0]], v[1])
			edgelen[v[0]+"-"+v[1]] = values[ind]
		}

		if _, ok := edgelen[v[1]+"-"+v[0]]; !ok {
			graph[v[1]] = append(graph[v[1]], v[0])
			edgelen[v[1]+"-"+v[0]] = 1.0 / values[ind]
		}

		if _, ok := edgelen[v[0]+"-"+v[0]]; !ok {
			edgelen[v[0]+"-"+v[0]] = 1.0
		}

		if _, ok := edgelen[v[1]+"-"+v[1]]; !ok {
			edgelen[v[1]+"-"+v[1]] = 1.0
		}
	}

	var results []float64

	for _, q := range queries {
		_, ok1 := graph[q[0]]
		_, ok2 := graph[q[1]]
		if !ok1 || !ok2 {
			results = append(results, -1.0)
			continue
		}
		if q[0] == q[1] {
			results = append(results, 1.0)
			continue
		}

		if _, ok := edgelen[q[0]+"-"+q[1]]; ok {
			results = append(results, edgelen[q[0]+"-"+q[1]])
			continue
		}
		if _, ok := edgelen[q[1]+"-"+q[0]]; ok {
			results = append(results, edgelen[q[1]+"-"+q[0]])
			continue
		}

		visited := make(map[string]bool, n)
		dist := dfs(q[0], q[1], graph, edgelen, visited)

		results = append(results, dist)
	}

	return results
}

func dfs(start, end string, graph map[string][]string, edgelen map[string]float64, visited map[string]bool) float64 {
	visited[start] = true

	nodes := graph[start]
	for _, n := range nodes {
		if visited[n] {
			continue
		}
		visited[n] = true
		if n == end {
			return edgelen[start+"-"+n]
		}

		d := dfs(n, end, graph, edgelen, visited)
		if d > 0.0 {
			return d * edgelen[start+"-"+n]
		}
	}

	return -1.0
}

func calcEquationV2(equations [][]string, values []float64, queries [][]string) []float64 {
	vertices := make(map[string]int)
	graph := make(map[int][]int)
	edgelen := make(map[int]float64)

	vertex := 1
	for ind, v := range equations {
		v0, ok1 := vertices[v[0]]
		v1, ok2 := vertices[v[1]]
		if !ok1 {
			vertices[v[0]], v0 = vertex, vertex
			vertex++
		}
		if !ok2 {
			vertices[v[1]], v1 = vertex, vertex
			vertex++
		}

		if _, ok := edgelen[v0*10000+v1]; !ok {
			graph[v0] = append(graph[v0], v1)
			edgelen[v0*10000+v1] = values[ind]
		}

		if _, ok := edgelen[v1*10000+v0]; !ok {
			graph[v1] = append(graph[v1], v0)
			edgelen[v1*10000+v0] = 1.0 / values[ind]
		}

		if _, ok := edgelen[v0*10000+v0]; !ok {
			edgelen[v0*10000+v0] = 1.0
		}

		if _, ok := edgelen[v1*10000+v1]; !ok {
			edgelen[v1*10000+v1] = 1.0
		}
	}

	var results []float64

	for _, q := range queries {
		q0, ok1 := vertices[q[0]]
		q1, ok2 := vertices[q[1]]
		if !ok1 || !ok2 {
			results = append(results, -1.0)
			continue
		}
		if q[0] == q[1] {
			results = append(results, 1.0)
			continue
		}

		if _, ok := edgelen[q0*10000+q1]; ok {
			results = append(results, edgelen[q0*10000+q1])
			continue
		}
		if _, ok := edgelen[q1*10000+q0]; ok {
			results = append(results, edgelen[q1*10000+q0])
			continue
		}

		visited := make(map[int]bool)
		dist := dfsV2(q0, q1, graph, edgelen, visited)

		results = append(results, dist)
	}

	return results
}

func dfsV2(start, end int, graph map[int][]int, edgelen map[int]float64, visited map[int]bool) float64 {
	visited[start] = true

	nodes := graph[start]
	for _, n := range nodes {
		if visited[n] {
			continue
		}
		visited[n] = true
		if n == end {
			return edgelen[start*10000+n]
		}

		d := dfsV2(n, end, graph, edgelen, visited)
		if d > 0.0 {
			return d * edgelen[start*10000+n]
		}
	}

	return -1.0
}

type DisjointSet struct {
	root  []int
	rank  []int
	value []float64
}

func (d *DisjointSet) Init() {
	d.root = make([]int, 0)
	d.rank = make([]int, 0)
	d.value = make([]float64, 0)
}

func (d *DisjointSet) find(x int) (int, float64) {
	if x == d.root[x] {
		return x, d.value[x]
	}
	var v float64
	d.root[x], v = d.find(d.root[x])
	d.value[x] *= v
	return d.root[x], d.value[x]
}

func (d *DisjointSet) union(x, y int, v float64) {
	rootX, _ := d.find(x)
	rootY, _ := d.find(y)
	if rootX == rootY {
		return
	}
	switch {
	case d.rank[rootX] > d.rank[rootY]:
		{
			d.root[rootY] = rootX
			d.value[rootY] = d.value[x] * v / d.value[y]
		}
	case d.rank[rootY] > d.rank[rootX]:
		{
			d.root[rootX] = rootY
			d.value[rootX] = d.value[y] / (d.value[x] * v)
		}
	default:
		{
			d.root[rootY] = rootX
			d.rank[rootX] += 1
			d.value[rootY] = d.value[x] * v / d.value[y]
		}
	}
}

func calcEquation_leet0ms(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]int)
	var ds DisjointSet
	ds.Init()
	for i, equation := range equations {
		if _, ok := m[equation[0]]; !ok {
			m[equation[0]] = len(ds.root)
			ds.root = append(ds.root, len(ds.root))
			ds.rank = append(ds.rank, 0)
			ds.value = append(ds.value, 1)
		}
		if _, ok := m[equation[1]]; !ok {
			m[equation[1]] = len(ds.root)
			ds.root = append(ds.root, len(ds.root))
			ds.rank = append(ds.rank, 0)
			ds.value = append(ds.value, 1)
		}
		ds.union(m[equation[0]], m[equation[1]], values[i])
	}
	ans := make([]float64, len(queries))
	for i, query := range queries {
		root1, ok1 := m[query[0]]
		root2, ok2 := m[query[1]]
		if ok1 && ok2 {
			r1, v1 := ds.find(root1)
			r2, v2 := ds.find(root2)
			if r1 == r2 {
				ans[i] = v2 / v1
				continue
			}
		}
		ans[i] = -1
	}
	return ans
}
