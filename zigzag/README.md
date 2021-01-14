#### 01/14/2021

```
~/leet/zigzag$ go test -run=XXX -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: leet/zigzag
BenchmarkConvertShort/my0msV1-8         	39311454	       139 ns/op	      40 B/op	       1 allocs/op
BenchmarkConvertShort/leet-4ms-8        	32139037	       187 ns/op	      62 B/op	       2 allocs/op
BenchmarkConvertLong/my0msV1-8          	 1379788	      3932 ns/op	    1690 B/op	       2 allocs/op
BenchmarkConvertLong/leet-4ms-8         	 1384233	      4346 ns/op	    2251 B/op	       7 allocs/op
PASS
ok  	leet/zigzag	32.058s
```

###### Benchmarks
`Short` has short input.
`Long` has long input.
