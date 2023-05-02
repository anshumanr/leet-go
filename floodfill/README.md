#### 05/01/2023

```
~/leet/floodfill$ go test  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/floodfill
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkFloodfill/mine-8               261233816                4.575 ns/op           0 B/op          0 allocs/op
BenchmarkFloodfill/minev2-8             261788968                4.473 ns/op           0 B/op          0 allocs/op
BenchmarkFloodfill/leet-fastest-8       259717329                4.720 ns/op           0 B/op          0 allocs/op
PASS
ok      leet/floodfill  4.999s
```