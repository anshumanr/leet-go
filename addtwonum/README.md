##### 05/06/2020

[Add two numbers](https://leetcode.com/problems/add-two-numbers)

```
~/code/addtwonum$ go test -run=XX -bench=.
goos: linux
goarch: amd64
BenchmarkAddTwoNumbers/addTwoNumbersV1-8                 1966377               598 ns/op             304 B/op         19 allocs/op
BenchmarkAddTwoNumbers/addTwoNumbersV2-8                  605988              1866 ns/op            1624 B/op         48 allocs/op
BenchmarkAddTwoNumbers/addTwoNumbersV3-8                 1000000              1029 ns/op             496 B/op         31 allocs/op
BenchmarkAddTwoNumbers/addTwoNumbers-0ms-8                981014              1046 ns/op             496 B/op         31 allocs/op
BenchmarkAddTwoNumbers/addTwoNumbers-4ms-8              41416851                25.9 ns/op             0 B/op          0 allocs/op
PASS
ok      _/home/anshuman/code/addtwonum  6.134s
```

`Note` - 4ms solution is clever clobbers the input