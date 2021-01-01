#### 01/01/2021

```

~/leet/3sumClosest$ go test -run=XXX -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: leet/3sumClosest
Benchmark3sum/my3sumClosest-8              35924            163460 ns/op              32 B/op          1 allocs/op
Benchmark3sum/threeSumClosest0ms-8         29038            212223 ns/op              32 B/op          1 allocs/op
PASS
ok      leet/3sumClosest        15.807s
```

`Note` - on leetcode, my version took 4ms