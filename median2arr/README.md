#### 01/06/2021

```
~/leet/median2arr$ go test -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: leet/median2arr
BenchmarkLongSubstr/myV1-static-8         	 2435014	      2476 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongSubstr/myV1-dynamic-8        	 1000000	      5173 ns/op	   16384 B/op	       1 allocs/op
BenchmarkLongSubstr/myV3-8                	 2300599	      2571 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongSubstr/myV4-8                	 4555452	      1314 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongSubstr/leet-4ms-8            	28292188	       204 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongSubstr/leet-8ms-8            	 1198574	      4937 ns/op	   16384 B/op	       1 allocs/op
PASS
ok  	leet/median2arr	46.315s
```