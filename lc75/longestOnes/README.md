### 05/07/2023

```
~/leet/lc75/longestOnes$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/longestOnes
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMaxOperations/mine-8             157081              6949 ns/op               0 B/op          0 allocs/op
BenchmarkMaxOperations/leet-fastest-8     165261              6608 ns/op               0 B/op          0 allocs/op
PASS
ok      leet/lc75/longestOnes   2.341s
```
```
~/leet/lc75/longestOnes$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/longestOnes
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMaxOperations/mine-8             152451              6864 ns/op               0 B/op          0 allocs/op
BenchmarkMaxOperations/leet-fastest-8     166366              6535 ns/op               0 B/op          0 allocs/op
PASS
ok      leet/lc75/longestOnes   2.293s
```
