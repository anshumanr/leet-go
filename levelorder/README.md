#### 04/26/2023

```
Results are very questionable as input size is very small
~/leet/levelorder$ go test  -bench=. -benchtime 5s -benchmem
goos: linux
goarch: amd64
pkg: leet/levelorder
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLevelOrder/mine-8              15704654               409.5 ns/op           128 B/op          5 allocs/op
BenchmarkLevelOrder/leet-1ms-8          16613252               314.5 ns/op           152 B/op          7 allocs/op
PASS
ok      leet/levelorder 12.420s
```