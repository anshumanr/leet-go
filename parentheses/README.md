#### 04/17/2023

```
~/leet/parentheses$ go test -bench=.
goos: linux
goarch: amd64
pkg: leet/parentheses
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLongestValidParentheses/mine-8                 19787588                54.16 ns/op           16 B/op          1 allocs/op
BenchmarkLongestValidParentheses/mine2-8                24165456                49.92 ns/op           36 B/op          1 allocs/op
BenchmarkLongestValidParentheses/leetcode-8             35203258                34.45 ns/op           36 B/op          1 allocs/op
PASS
ok      leet/parentheses        5.488s
```
