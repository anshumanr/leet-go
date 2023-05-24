#### 05/24/2023

```
## too many allocations, and searching
~/leet/lc75/binarytree/pathsum$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/binarytree/pathsum
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkGoodNodes/mine-8                  26168             46865 ns/op           57516 B/op       1182 allocs/op
BenchmarkGoodNodes/leet-fastest-8          68991             16511 ns/op             684 B/op          1 allocs/op
PASS
ok      leet/lc75/binarytree/pathsum    3.017s
```