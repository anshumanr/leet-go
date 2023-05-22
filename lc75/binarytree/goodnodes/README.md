#### 05/22/2023

```
## both variants are same
~/leet/lc75/binarytree/goodnodes$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/binarytree/goodnodes
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkGoodNodes/mine-8               41230086                29.15 ns/op            0 B/op          0 allocs/op
BenchmarkGoodNodes/leet-fastest-8       23636274                44.53 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/binarytree/goodnodes  3.306s
```