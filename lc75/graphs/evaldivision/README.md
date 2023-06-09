#### 06/09/2023

```
### Need to learn disjoint sets
~/leet/lc75/graphs/evaldivision$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/graphs/evaldivision
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkCalcEquation/mine-8              337488              3579 ns/op             562 B/op         19 allocs/op
BenchmarkCalcEquation/mine-v2-8           554098              1807 ns/op             377 B/op         10 allocs/op
BenchmarkCalcEquation/leet-fastest-8     1267173               880.8 ns/op           312 B/op         11 allocs/op
PASS
ok      leet/lc75/graphs/evaldivision   5.308s
anshuman@r-19-de-22-695l:~/leet/lc75/graphs/evaldivision$
```
