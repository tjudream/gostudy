# 37 BDD
## Behavior Driven Development
* 让业务领域的专家参与开发
* 你知道Story Card背面应该写什么吗？
    * 背面应该写该Stroy应该如何被验收
* 用业务领域的语言来描述
    * "Given - When - Then"
    * <b><u>Given</u></b> a user is creating an account
    * <b><u>When</u></b> they specify an insecure password
    * <b><u>Then</u></b> they see a message, "Passwords must be at least 6 characters long with at least one letter,
    one number, and one symbol."
## BDD in Go
* 项目网站： [https://github.com/smartystreets/goconvey](https://github.com/smartystreets/goconvey)
* 安装
```go
go get -u github.com/smartystreets/goconvey/convey
```
* 启动WEB UI
```go
$GOPATH/bin/goconvey
```
