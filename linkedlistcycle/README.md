#### 04/20/2023

```
~/leet/linkedlistcycle$ go test -run=XXX -bench=.
goos: linux
goarch: amd64
pkg: leet/linkedlistcycle
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkLinkedListCycle/mine-8                 13432845                82.61 ns/op            0 B/op          0 allocs/op
BenchmarkLinkedListCycle/mine-fp-8              297500455                4.003 ns/op           0 B/op          0 allocs/op
BenchmarkLinkedListCycle/leet-0ms-8             306626217                3.867 ns/op           0 B/op          0 allocs/op
PASS
ok      leet/linkedlistcycle    4.392s
```