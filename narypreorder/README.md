#### 04/20/2023

```
Results are very questionable as input size is very small
~/leet/narypreorder$ go test  -bench=. -benchtime 5s -benchmem
goos: linux
goarch: amd64
pkg: leet/narypreorder
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkNaryPreorder/mine-8            13064130               452.0 ns/op           372 B/op         18 allocs/op
BenchmarkNaryPreorder/mine-iter-8         397812             13500 ns/op           82104 B/op          5 allocs/op
BenchmarkNaryPreorder/leet-0ms-8        15242287               396.6 ns/op           280 B/op         10 allocs/op
PASS
ok      leet/narypreorder       18.354s
```