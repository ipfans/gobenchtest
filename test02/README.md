# 类型判断和接口性能效率测试

```
➜  test02 git:(master) go test -v -bench=. -benchmem
testing: warning: no tests to run
Benchmark_TypeSwitch-4        	50000000	        30.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_NormalSwitch-4      	1000000000	         2.11 ns/op	       0 B/op	       0 allocs/op
Benchmark_InterfaceSwitch-4   	100000000	        18.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ipfans/gobenchtest/test02	5.769s
```