# interface变量性能

```
➜  test03 git:(master) ✗ go test -v -bench=. -benchmem
testing: warning: no tests to run
Benchmark_Normal-4         	2000000000	         0.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_NormalRef-4      	2000000000	         0.40 ns/op	       0 B/op	       0 allocs/op
Benchmark_Interface-4      	100000000	        23.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_InterfaceRef-4   	2000000000	         0.45 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ipfans/gobenchtest/test03	5.003s
```