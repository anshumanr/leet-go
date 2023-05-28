#### 05/28/2023

```
### V2 is inspired by leet/fastest
~/leet/lc75/binarytree/zigzag$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/binarytree/zigzag
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLongestZigZag/mine-8            1029952              1208 ns/op               0 B/op          0 allocs/op
BenchmarkLongestZigZag/mine-V2-8         1409074               883.4 ns/op             0 B/op          0 allocs/op
BenchmarkLongestZigZag/leet-fastest-8    1342294               978.4 ns/op             8 B/op          1 allocs/op
PASS
ok      leet/lc75/binarytree/zigzag     6.452s
```