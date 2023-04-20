#### 04/20/2023

```
~/leet/reverselist$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: leet/reverselist
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMerge2Lists/mine-8             334150807                3.541 ns/op           0 B/op          0 allocs/op
BenchmarkMerge2Lists/leet-0ms-8         18595104                57.47 ns/op           24 B/op          2 allocs/op
BenchmarkMerge2Lists/leet-1ms-8         281657929                4.114 ns/op           0 B/op          0 allocs/op
PASS
ok      leet/reverselist        4.284s
```