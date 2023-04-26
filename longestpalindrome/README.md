#### 04/26/2023

```
~/leet/longestpalindrome$ go test  -bench=. -benchtime 5s -benchmem
goos: linux
goarch: amd64
pkg: leet/longestpalindrome
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLongestPalindrome/mine-8                5373733              1111 ns/op               0 B/op          0 allocs/op
BenchmarkLongestPalindrome/leet-0ms-8             418852             12709 ns/op             154 B/op          2 allocs/op
PASS
ok      leet/longestpalindrome  12.575s
```
