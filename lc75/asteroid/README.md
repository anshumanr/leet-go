#### 05/05/2023

```
~/leet/lc75/asteroid$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/asteroid
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkAsteroidCollision/mine-8                 786249              1522 ns/op            3429 B/op          1 allocs/op
BenchmarkAsteroidCollision/leet-fastest-8         787878              1326 ns/op             352 B/op          3 allocs/op
PASS
ok      leet/lc75/asteroid      3.267s
```
