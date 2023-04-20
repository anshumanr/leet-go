#### 04/20/2023

```
~/leet/middlelist$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: leet/middlelist
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkMiddleNode/mine-8              60685917                17.11 ns/op            0 B/op          0 allocs/op
BenchmarkMiddleNode/mine-fp-8           281915134                4.217 ns/op           0 B/op          0 allocs/op
BenchmarkMiddleNode/leet-1ms-8           8590219               133.0 ns/op            69 B/op          2 allocs/op
PASS
ok      leet/middlelist 3.971s
```