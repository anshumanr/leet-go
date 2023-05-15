package asteroid

func asteroidCollision(asteroids []int) []int {
	l := len(asteroids)

	stack := make([]int, l)
	top := 0

	for i := 0; i < l; i++ {
		if top > 0 && stack[top-1] > 0 && asteroids[i] < 0 {
			abs := asteroids[i] * -1
			if stack[top-1] > abs {
				continue
			}
			if stack[top-1] == abs {
				top--
			} else {
				top--
				i--
			}
		} else {
			stack[top] = asteroids[i]
			top++
		}
	}

	return stack[:top]
}

func asteroidCollision_leet2ms(asteroids []int) []int {
	stack := make([]int, 0)

	for i := range asteroids {
		if len(stack) > 0 && stack[len(stack)-1] > 0 && asteroids[i] < 0 {
			destroyed := false
			for len(stack) > 0 && !destroyed && stack[len(stack)-1] > 0 {
				if abs(stack[len(stack)-1]) > abs(asteroids[i]) {
					destroyed = true
				} else if abs(stack[len(stack)-1]) < abs(asteroids[i]) {
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					destroyed = true
				}
			}

			if !destroyed {
				stack = append(stack, asteroids[i])
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}

	return stack
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
