#### 05/07/2023

```
~/leet/lc75/removestars$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/removestars
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkReverseVowels/mine-8             677721              1839 ns/op             855 B/op          1 allocs/op
BenchmarkReverseVowels/using-stack-8     1980273               596.6 ns/op           864 B/op          1 allocs/op
BenchmarkReverseVowels/leet-fastest-8    1336232               904.8 ns/op          1455 B/op          4 allocs/op
PASS
ok      leet/lc75/removestars   6.101s
```
