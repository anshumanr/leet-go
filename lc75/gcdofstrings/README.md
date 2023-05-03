#### 05/03/2023

```
//Note: slicing operations seem way more expensive than simple looping & recursing
~/leet/lc75/gcdOfStrings$ go test  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/gcdofstrings
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkGcdOfStrings/mine-8            10535973               107.8 ns/op             0 B/op          0 allocs/op
BenchmarkGcdOfStrings/mine-v2-8         13347388                76.81 ns/op            0 B/op          0 allocs/op
BenchmarkGcdOfStrings/leet-fastest-8    21353858                50.43 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/gcdofstrings  3.511s
```
