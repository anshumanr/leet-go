package smallestnum

type SmallestInfiniteSet struct {
	arr []bool
	min int
}

func Constructor() SmallestInfiniteSet {
	a := make([]bool, 1001)
	for i := range a {
		a[i] = true
	}

	return SmallestInfiniteSet{
		arr: a,
		min: 0,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i, v := range this.arr {
		if v == true {
			this.arr[i] = false
			return i + 1
		}
	}

	return -1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.arr[num-1] = true
	if num < this.min {
		this.min = num
	}
}
