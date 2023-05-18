#### 05/16/2023

```
~/leet/lc75/decodestring$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/decodestring
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkReverseVowels/mine-8             383574              3090 ns/op            9845 B/op         19 allocs/op
BenchmarkReverseVowels/leet-fastest-8     483507              2470 ns/op            8089 B/op         15 allocs/op
PASS
ok      leet/lc75/decodestring  3.963s
```
