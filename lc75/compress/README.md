#### 05/06/2023

```
~/leet/lc75/compress$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/compress
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkCompress/mine-8                 1000000              1028 ns/op               0 B/op          0 allocs/op
BenchmarkCompress/leet-fastest-8          746205              1668 ns/op               0 B/op          0 allocs/op
PASS
ok      leet/lc75/compress      3.290s
```
