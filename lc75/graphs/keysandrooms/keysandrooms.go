package keysandrooms

func canVisitAllRooms(rooms [][]int) bool {
	l := len(rooms)
	v := make(map[int]bool)

	var keys []int
	v[0] = true
	keys = append(keys, 0)

	j := len(keys)
	for j > 0 {
		k := keys[j-1]
		keys = keys[:j-1]
		for _, rk := range rooms[k] {
			if _, ok := v[rk]; !ok {
				v[rk] = true
				keys = append(keys, rk)
			}
		}

		j = len(keys)
	}

	return len(v) == l
}

// Goal: Check if it is possible to visit all rooms

// Conditions: All rooms are locked except of room 0

// Approach: Use bfs
// Go to first roome
// grab keys and put them into the queue
// if you visit a room put set room to true
// At the end iterate to all rooms and check if all rooms are visited

func canVisitAllRooms_leet0ms(rooms [][]int) bool {
	visitedRooms := make([]bool, len(rooms))
	queue := []int{0}

	for len(queue) > 0 {
		currentRoom := queue[0]
		queue = queue[1:]
		if len(rooms[currentRoom]) == 0 || visitedRooms[currentRoom] {
			visitedRooms[currentRoom] = true
			continue
		} else {
			visitedRooms[currentRoom] = true
		}
		queue = append(queue, rooms[currentRoom]...)
	}

	for _, visited := range visitedRooms {
		if !visited {
			return false
		}
	}

	return true
}
