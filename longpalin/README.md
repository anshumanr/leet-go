#### 01/11/2021

```
/leet/longpalin$ go test -run=XXX -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: leet/longpalin
BenchmarkLongPalSame/myV3-8         	    7892	    706774 ns/op	       4 B/op	       1 allocs/op
BenchmarkLongPalSame/leet-0ms-8     	 8763310	       677 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongPalSame/leet-4ms-8     	    7146	    825476 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongPalRandom/myV3-8       	 3555732	      1693 ns/op	       4 B/op	       1 allocs/op
BenchmarkLongPalRandom/leet-0ms-8   	10932316	       541 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongPalRandom/leet-4ms-8   	 2647800	      2290 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongPalLen1/myV3-8         	  916086	      6221 ns/op	       4 B/op	       1 allocs/op
BenchmarkLongPalLen1/leet-0ms-8     	 2908334	      2028 ns/op	       0 B/op	       0 allocs/op
BenchmarkLongPalLen1/leet-4ms-8     	  672504	      8501 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	leet/longpalin	60.386s
```

###### Benchmarks
`Same` as same char repeated.
`Random` has mixed input.
`Len1` has long input with longest palindrome length 1.

`V1` implementation was incorrect, `V2` timed out (very slow), `V3` is faster than `4ms` but `0ms` is clearly the best.