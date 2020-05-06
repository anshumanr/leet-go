##### 05/05/2020

```
~/code/two-sum$ go test -bench=. -benchtime=2s
goos: linux
goarch: amd64
pkg: twoSum
BenchmarkTwoSum/twoSum-0ms-8             3789748               631 ns/op            1168 B/op          2 allocs/op
BenchmarkTwoSum/twoSum-4ms-8             9593313               252 ns/op              16 B/op          1 allocs/op
BenchmarkTwoSum/twoSum-8ms-8             9204481               254 ns/op              16 B/op          1 allocs/op
BenchmarkTwoSum/twoSum-12ms-8             827277              2515 ns/op             960 B/op          4 allocs/op
BenchmarkTwoSum/twoSum-20ms-8           58091673                39.8 ns/op            16 B/op          1 allocs/op
BenchmarkTwoSum/twoSum-32ms-8           51000418                46.9 ns/op            16 B/op          1 allocs/op
BenchmarkTwoSum/twoSumAttempt1-8        57569691                40.7 ns/op            16 B/op          1 allocs/op
BenchmarkTwoSum/twoSumAttempt2-8        61128237                38.7 ns/op            16 B/op          1 allocs/op
BenchmarkTwoSum/twoSumAttempt3-8        56503388                40.3 ns/op            16 B/op          1 allocs/op
PASS
ok      twoSum  22.340s
```

Benchmarks do not agree with leet code results
