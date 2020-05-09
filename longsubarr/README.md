#### 05/08/2020

Benchmark 1: input size: 10^5, diff: 6466408

```
~/code/longsubarr$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: longsubarr
BenchmarkAtoi/longestSubarrayV3-8                    842           1401515 ns/op               0 B/op          0 allocs/op
BenchmarkAtoi/longestSubarray-8                       61          16718279 ns/op        11200124 B/op     400000 allocs/op
PASS
ok      longsubarr      2.368s
```


Benchmark 1: input size: 10^5, diff: 1000000000
```
~/code/longsubarr$ 
anshuman@anshuman-Inspiron-15:~/code/longsubarr$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: longsubarr
BenchmarkAtoi/longestSubarrayV3-8                   2911            383784 ns/op               0 B/op          0 allocs/op
BenchmarkAtoi/longestSubarray-8                       70          16721895 ns/op        11200118 B/op     400000 allocs/op
PASS
ok      longsubarr      3.328s
```

`Note:` longestSubarray-8 is solution from leet code