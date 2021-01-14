package zigzag

func convert(s string, numRows int) string {
	l := len(s)
	grid := make([]byte, l)
	adj := numRows - 2
	if adj < 0 {
		adj = 0
	}
	slot := 0

	inc1, inc2, cnt := 0, 0, 0
	for cnt < numRows {
		inc1 = (numRows + adj) - 2*cnt
		inc2 = (numRows + adj) - inc1

		t := true
		for i := cnt; i < l; {
			grid[slot] = s[i]
			slot++

			//first row
			if inc2 == 0 {
				i += inc1
				continue
			}
			if inc1 == 0 {
				i += inc2
				continue
			}

			if t {
				i += inc1
				t = false
			} else {
				i += inc2
				t = true
			}
		}
		cnt++
	}

	return string(grid)
}

type iterator struct {
	pos    int // 0..length-1
	depth  int // 0..rows
	step   int // 0,2,4,6,...
	delta  int // const
	rows   int // const
	length int // const
}

func convertLeet4ms(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	it := iterator{pos: 0, depth: 0, step: (numRows - 1) * 2, delta: (numRows - 1) * 2, rows: numRows, length: len(s)}
	var z []byte
	for true {
		z = append(z, s[it.pos])
		if it.next() == false {
			break
		}
	}

	return string(z)
}

func (it *iterator) next() bool {
	it.pos += it.step
	if it.step != it.delta {
		it.step = it.delta - it.step
	}

	if it.pos < it.length {
		return true
	}

	it.depth++
	if it.depth == it.rows {
		return false
	}

	it.pos = it.depth
	it.step = (it.rows - it.depth - 1) * 2
	if it.step == 0 {
		it.step = it.delta
	}

	if it.pos < it.length {
		return true
	}

	return false
}
