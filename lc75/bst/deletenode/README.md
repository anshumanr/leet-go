#### 06/04/2023

```
~/leet/lc75/bst/deletenode$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/bst/deletenode
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkDeleteNode/mine-8              59249301                18.60 ns/op            0 B/op          0 allocs/op
BenchmarkDeleteNode/leet-fastest-8      70702654                17.03 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/bst/deletenode        3.266s
```