#### 05/07/2023

Note: v1 was brute force using `strings.Count`
```
~/leet/lc75/maxvowels$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/maxvowels
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkCompress/mine-8                     100          27039940 ns/op               0 B/op          0 allocs/op
BenchmarkCompress/mine-v3-8                19808             58944 ns/op               0 B/op          0 allocs/op
BenchmarkCompress/leet-fastest-8           20188             57179 ns/op               0 B/op          0 allocs/op
PASS
ok      leet/lc75/maxvowels     6.250s
```
