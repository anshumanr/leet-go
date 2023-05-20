#### 05/20/2023

```
## Repeat: both variants are same
~/leet/lc75/binarytree/lca$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/binarytree/lca
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLCA/mine-8             58859083                20.26 ns/op            0 B/op          0 allocs/op
BenchmarkLCA/leet-fastest-8     59475134                20.13 ns/op            0 B/op          0 allocs/op
PASS
ok      leet/lc75/binarytree/lca        4.334s
```