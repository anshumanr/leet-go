### 05/12/2023

Note: Difference between original & v2: 1. got rid of slice `visited` 2. used `[]pos` instead of `[][]int` to queue position in grid.

```
~/leet/lc75/graphs/exitmaze$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/graphs/exitmaze
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkNearestExit/mine-8              2035440               576.0 ns/op           312 B/op         10 allocs/op
BenchmarkNearestExit/mine-v2-8          22213654                48.96 ns/op           16 B/op          1 allocs/op
BenchmarkNearestExit/leet-fastest-8     35691871                29.17 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/graphs/exitmaze       3.998s
```
