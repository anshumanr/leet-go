package placeflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	if l == 1 {
		if flowerbed[0] == 0 {
			return n <= 1
		}
		return n == 0
	}

	cnt := 0
	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		cnt++
	}

	for i := 1; i < l-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			cnt++
		}
	}

	if flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
		flowerbed[l-1] = 1
		cnt++
	}

	return cnt >= n
}

func canPlaceFlowers_leet2ms(flowerbed []int, n int) bool {
	size := len(flowerbed)

	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == size-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			i++
		}
	}

	return n <= 0
}
