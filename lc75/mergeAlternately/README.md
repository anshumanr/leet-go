#### 04/26/2023

```
~/leet/lc75/mergeAlternately$ go test  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/mergeAlternately
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMergeAlternately/mine-8                22020937                47.12 ns/op           16 B/op          2 allocs/op
BenchmarkMergeAlternately/leet-fastest-8        24738609                48.64 ns/op           16 B/op          2 allocs/op
PASS
ok      leet/lc75/mergeAlternately      3.305s
```
