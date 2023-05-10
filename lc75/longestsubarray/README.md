### 05/07/2023

Using `if` to compare & sway for longest in V2 (see commented code).
```
~/leet/lc75/longestsubarray$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/longestsubarray
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMaxArea/mine-8                   889004              1295 ns/op               0 B/op          0 allocs/op
BenchmarkMaxArea/mine-V2-8               1683080               699.2 ns/op             0 B/op          0 allocs/op
BenchmarkMaxArea/leet-fastest-8          2000937               591.1 ns/op             0 B/op          0 allocs/op
PASS
ok      leet/lc75/longestsubarray       4.861s
```
Using `max()` to compare & swap in V2.
```
~/leet/lc75/longestsubarray$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/longestsubarray
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMaxArea/mine-8                   844869              1326 ns/op               0 B/op          0 allocs/op
BenchmarkMaxArea/mine-V2-8               1979499               598.6 ns/op             0 B/op          0 allocs/op
BenchmarkMaxArea/leet-fastest-8          2012508               596.3 ns/op             0 B/op          0 allocs/op
PASS
ok      leet/lc75/longestsubarray       4.747s
```
