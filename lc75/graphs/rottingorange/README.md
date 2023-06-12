### 05/12/2023

```
$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/graphs/rottingorange
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkOrangesRotting/mine-8           5598958               207.3 ns/op           276 B/op          3 allocs/op
BenchmarkOrangesRotting/leet-fastest-8   5506402               199.2 ns/op           184 B/op          3 allocs/op
PASS
ok      leet/lc75/graphs/rottingorange  2.705s
```
