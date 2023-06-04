#### 06/04/2023

```
### Too small an input to be reliable
~/leet/lc75/graphs/keysandrooms$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/graphs/keysandrooms
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkCanVisitAllRooms/mine-8                 8948683               121.9 ns/op            16 B/op          1 allocs/op
BenchmarkCanVisitAllRooms/leet-fastest-8         9978061               116.7 ns/op            64 B/op          4 allocs/op
PASS
ok      leet/lc75/graphs/keysandrooms   2.524s
```
