#### 06/04/2023

```
$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/graphs/reorderroutes
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMinReorder/mine-8                   921           1132648 ns/op          409114 B/op      16621 allocs/op
BenchmarkMinReorder/leet-fastest-8           921           1122119 ns/op          409111 B/op      16621 allocs/op
PASS
ok      leet/lc75/graphs/reorderroutes  2.346s
```
