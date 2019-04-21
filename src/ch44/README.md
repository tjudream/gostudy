# 44 easyjson
## 更快的 JSON 解析
* EasyJSON 采用代码生成而非反射
```go
    goos: darwin
    goarch: amd64
    BenchmarkEmbeddedJson-4         200000              6360 ns/op
    BenchmarkEasyJson-4            1000000              1396 ns/op
```
* 安装
```jshelllanguage
    go get -u github.com/mailru/easyjson/...
```
* 使用
```jshelllanguage
    easyjson -all <结构定义>.go
```