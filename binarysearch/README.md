#### 04/27/2023

```
~/leet/binarysearch$ go test  -bench=. -benchtime 5s -benchmem
goos: linux
goarch: amd64
pkg: leet/binarysearch
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkSearch/mine-8          525254850               11.36 ns/op            0 B/op          0 allocs/op
BenchmarkSearch/mine-v2-8       490146403               12.06 ns/op            0 B/op          0 allocs/op
BenchmarkSearch/leet-17ms-8     626877529                9.566 ns/op           0 B/op          0 allocs/op
PASS
ok      leet/binarysearch       21.251s
```

##### Note:
Difference between v2 & leet code is the where mid is being declared & defined.