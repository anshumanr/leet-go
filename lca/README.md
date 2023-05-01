#### 05/01/2023

```
### Learnt from the really bad result!!
~/leet/lca$ go test  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lca
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLCA/mine-8             230951431                5.268 ns/op           0 B/op          0 allocs/op
BenchmarkLCA/leet-fastest-8     256370522                4.747 ns/op           0 B/op          0 allocs/op
PASS
ok      leet/lca        3.443s
```