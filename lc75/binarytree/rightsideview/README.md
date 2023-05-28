#### 05/28/2023

```
~/leet/lc75/binarytree/rightsideview$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/binarytree/rightsideview
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRightSideView/mine-8            5657690               201.9 ns/op           100 B/op          2 allocs/op
BenchmarkRightSideView/leet-fastest-8    2041954               573.4 ns/op           521 B/op          9 allocs/op
PASS
ok      leet/lc75/binarytree/rightsideview      3.134s
```