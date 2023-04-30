#### 04/26/2023

```
### Really bad result!!
~/leet/validbst$ go test  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/validbst
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLevelOrder/mine-8                318187              3256 ns/op            1060 B/op         12 allocs/op
BenchmarkLevelOrder/leet-0ms-8          84401569                13.24 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/validbst   2.215s
```