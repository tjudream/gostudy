# 47 性能分析工具

## 准备工作
* 安装 graphviz
    * brew install graphviz
* 将 $GOPATH/bin 加入到 $PATH
    * Mac OS: 在 .bash_profile 中修改路径
* 安装 go-torch
    * go get -u github.com/uber/go-torch
    * 下载并复制 flamegraph.pl ([https://github.com/brendangregg/FlameGraph](https://github.com/brendangregg/FlameGraph)) 
    至 $GOPATH/bin 路径下
    * 将 $GOPATH/bin 加入 $PATH