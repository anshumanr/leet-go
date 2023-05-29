#### 05/28/2023

```
~/leet/lc75/binarytree/maxlevelsum$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/binarytree/maxlevelsum
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMaxLevelSum/mine-8               953373              1099 ns/op             108 B/op          3 allocs/op
BenchmarkMaxLevelSum/mine-V1-8            467073              2459 ns/op            4465 B/op         19 allocs/op
BenchmarkMaxLevelSum/leet-fastest-8       499056              2384 ns/op            4465 B/op         19 allocs/op
PASS
ok      leet/lc75/binarytree/maxlevelsum        5.403s
```