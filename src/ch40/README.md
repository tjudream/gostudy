# 40 不安全编程
## “不安全”行为的危险性
```go
    i := 10
    f := *(*float64)(unsafe.Pointer(&i))
```