###### 05/05/2020

```
~/code/atoi$ go test -bench=.
goos: linux
goarch: amd64
pkg: atoi
BenchmarkAtoi/myAtoi-1-8                 5091895               215 ns/op              32 B/op          1 allocs/op
BenchmarkAtoi/myAtoi-2-8                 7556982               156 ns/op               0 B/op          0 allocs/op
BenchmarkAtoi/myAtoi-3-8                52163636                20.9 ns/op             0 B/op          0 allocs/op
BenchmarkAtoi/atoiLeetCode-0ms-8        36514119                30.7 ns/op             0 B/op          0 allocs/op
PASS
ok      atoi    4.959s
```

`Note` - getting rid of map makes it faster than leet code