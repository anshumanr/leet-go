#### 12/31/2020

```

~/leet/3sum$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: leet/3sum
Benchmark3sum/my3sum-8                 1        1536846974 ns/op        1599524896 B/op 49985054 allocs/op
Benchmark3sum/my3sumV2-8           14548             82630 ns/op              96 B/op          3 allocs/op
Benchmark3sum/3sumLeetCode-24ms-8                  14672             83455 ns/op              96 B/op          3 allocs/op
PASS
ok      leet/3sum       5.633s
```

`Note`
1. `append` was a bottleneck for large inputs (tests timed out)
2. Major difference between V1 (append to result & remove duplicates later) v/s V2 (do not process duplicate input)