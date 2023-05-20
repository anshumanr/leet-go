#### 05/20/2023

```
## simple v/s fast pointer
~/leet/lc75/maxtwinsumlist$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/maxtwinsumlist
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkPairSum/mine-8                  6354571               177.8 ns/op           304 B/op          4 allocs/op
BenchmarkPairSum/leet-fastest-8         31108870                37.67 ns/op           16 B/op          1 allocs/op
PASS
ok      leet/lc75/maxtwinsumlist        3.526s
```