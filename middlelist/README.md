#### 04/20/2023

```
~/leet/middlelist$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: leet/middlelist
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMerge2Lists/mine-8             57088138                18.94 ns/op            0 B/op          0 allocs/op
BenchmarkMerge2Lists/leet-1ms-8          8665527               133.2 ns/op            69 B/op          2 allocs/op
PASS
ok      leet/middlelist 2.406s
```