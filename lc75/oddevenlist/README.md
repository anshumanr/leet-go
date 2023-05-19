#### 05/19/2023

```
~/leet/lc75/oddevenlist$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/oddevenlist
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkOddEvenList/mine-8             76088946                14.42 ns/op            0 B/op          0 allocs/op
BenchmarkOddEvenList/leet-fastest-8     74103495                13.98 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/oddevenlist   2.174s
```