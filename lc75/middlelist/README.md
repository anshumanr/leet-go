#### 04/20/2023

```
~/leet/lc75/middlelist$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/middlelist
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMerge2Lists/mine-8             108553198               10.98 ns/op            0 B/op          0 allocs/op
BenchmarkMerge2Lists/mine-V1-8          77888337                13.55 ns/op            0 B/op          0 allocs/op
BenchmarkMerge2Lists/leet-fastest-8     100000000               10.97 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/middlelist    4.260s
```