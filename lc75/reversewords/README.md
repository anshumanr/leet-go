#### 05/04/2023

```
//Note: HUGE difference in using slices & using string builder
~/leet/lc75/reversewords$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/reversewords
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkGcdOfStrings/mine-8                5754            207569 ns/op         1274070 B/op        405 allocs/op
BenchmarkGcdOfStrings/mine-V2-8            86935             13293 ns/op            3149 B/op          1 allocs/op
BenchmarkGcdOfStrings/leet-fastest-8       98473             11459 ns/op            8599 B/op          2 allocs/op
PASS
ok      leet/lc75/reversewords  4.673s
```
