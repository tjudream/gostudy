# 36 Benchmark
## Benchmark
* 文件需要以_test结尾，如 bench_test.go
* func要以Benchmark开头
* 函数的参数需要是 *testing.B
```go
func BenchmarkConcatStringByAdd(b *testing.B) {
	//与性能测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
	}
	b.StopTimer()
	//与性能测试无关的代码
}
```
* 在命令行中运行benchmark
```go
go test -bench=.
```

* bench的值：
    * . 表示运行所有benchmark函数
    * 也可以直接写benchmark函数的名字，指定要运行benchmark的函数,如: bench=BenchmarkConcatStringByAdd

* 研究代码块有多少次内存分配：
```go
go test -bench=. -benchmem
```
* -bench=<相关benchmark测试>
* Windows 下使用go test命令行时， -bench=. 应写为 -bench="."

