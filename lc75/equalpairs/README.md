### 05/13/2023

Simple & elegant v/s BF (mine)
```
~/leet/lc75/equalpairs$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/equalpairs
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkEqualPairs/mine-8                 15511             76133 ns/op           14609 B/op         22 allocs/op
BenchmarkEqualPairs/leet-fastest-8        377460              3167 ns/op             593 B/op          0 allocs/op
PASS
ok      leet/lc75/equalpairs    4.162s
```
