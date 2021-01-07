#### 01/06/2021

```
~/leet/mergeklist$ go test -run=XXX -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: leet/mergeklist
BenchmarkMergeKLists/myV1-8         	   52398	    114549 ns/op	   29784 B/op	     847 allocs/op
BenchmarkMergeKLists/leet-0ms-8     	   57643	     91456 ns/op	   29840 B/op	     858 allocs/op
BenchmarkMergeKLists/leet-5ms-8     	   51346	    118522 ns/op	   39632 B/op	      11 allocs/op
PASS
ok  	leet/mergeklist	20.756s
```