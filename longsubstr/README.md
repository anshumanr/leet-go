#### 01/01/2021

```
~/leet/longsubstr$ go test -run=XXX -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: leet/longsubstr
BenchmarkLongSubstr/myV1-8         	    4119	   1435916 ns/op	  778433 B/op	   24326 allocs/op
BenchmarkLongSubstr/myV2-8         	  135264	     40224 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongSubstr/leet-0ms-8     	   86505	     69787 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	leet/longsubstr	18.697s
```

`Note:` V2 works in a single `for` loop [O(n)]