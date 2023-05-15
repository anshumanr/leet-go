#### 05/15/2023

```
~/leet/lc75/uniqueoccur$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/uniqueoccur
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkProductExceptSelf/mine-8                 974948              1196 ns/op              48 B/op          2 allocs/op
BenchmarkProductExceptSelf/leet-fastest-8         747560              1588 ns/op              48 B/op          2 allocs/op
PASS
ok      leet/lc75/uniqueoccur   4.183s
```
