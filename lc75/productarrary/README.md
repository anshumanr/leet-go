#### 05/03/2023

```
//Note: unsurprisingly, memory allocations is the reason for being slow
~/leet/lc75/productarrary$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/productarrary
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkProductExceptSelf/mine-8                9664722               108.7 ns/op           176 B/op          3 allocs/op
BenchmarkProductExceptSelf/leet-fastest-8       22519869                52.98 ns/op           58 B/op          1 allocs/op
PASS
ok      leet/lc75/productarrary 3.402s
```
