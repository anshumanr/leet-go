#### 05/12/2023

```
~/leet/lc75/closestrings$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/closestrings
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkCloseStrings/mine-8             3032424               385.4 ns/op           440 B/op          3 allocs/op
BenchmarkCloseStrings/leet-fastest-8     3068055               382.4 ns/op           440 B/op          3 allocs/op
PASS
ok      leet/lc75/closestrings  3.155s
```
