#### 05/16/2023

```
~/leet/lc75/dota2$ go test -run=XXX  -bench=. -benchmem
goos: linux
goarch: amd64
pkg: leet/lc75/dota2
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkPredictPartyVictory/mine-8                 2869            375371 ns/op           10240 B/op          1 allocs/op
BenchmarkPredictPartyVictory/leet-fastest-8        19357             60263 ns/op          163840 B/op          2 allocs/op
PASS
ok      leet/lc75/dota2 2.923s
```
