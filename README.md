# 字符串拼接测试

fmt.Sprintf等等会带来大量的性能损失，在热路径上尽量避免使用。

```
➜  test01 git:(master) ✗ go test -v -bench=. -benchmem
testing: warning: no tests to run
BenchmarkSimpleStringAppend-4   	50000000	        29.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkSimpleStringFormat-4   	 5000000	       271 ns/op	      80 B/op	       5 allocs/op
PASS
ok  	github.com/ipfans/gobenchtest/test01	3.191s
```