# 21 构件可复用的模块（包）
* package
    1. 基本复用模块单元
        * 以首字母大写来表明可被包外代码访问
    2. 代码的package可以和所在的目录不一致
    3. 同一个目录里的 Go 代码的 package 要保持一致
* init 
    * 在 main 被执行前，所有依赖的 package 的 init 方法都会被执行
    * 不同包的 init 函数按照包导入的依赖关系决定执行顺序
    * 每个包可以有多个 init 函数
    * 包的每个源文件也可以有多个 init 函数，这点比较特殊
* 获取远程包
    * 获取远程包
        * go get github.com/easierway/concurrent_map
    * 更新远程包（第一次获取也可以这样用）
        * go get -u github.com/easierway/concurrent_map 
1. 通过go get来获取远程依赖
    * go get -u 强制从网络更新远程依赖
2. 注意代码在GitHub上的组织形式，以适应 go get
    * 直接以代码路径开始，不要有src,直接用src的下一层
    * 例如：[https://github.com/easierway/concurrent_map](https://github.com/easierway/concurrent_map)
    * 下载之后的路径： $GOPATH/src/github.com/easierway/concurrent_map/*.go

# 22 依赖管理
* Go 未解决的依赖问题
    1. 同一环境下，不同项目使用同一包的不同版本
    2. 无法管理对包的特定版本的依赖
* vendor 路径
   Go 1.5 release 版本，vendor目录被添加到了除了GOPATH和GOROOT之外的依赖目录查找的解决方案。在Go 1.6之前，需要手动的设置环境变量
* 查找依赖包路径的解决方案如下：
    1. 当前包下的vendor目录
    2. 向上级目录查找，直到找到src下的vendor目录
    3. 在GOPATH下面查找依赖包
    4. 在GOROOT目录下查找
* 常用的依赖管理工具
    * godep [https://github.com/tools/godep](https://github.com/tools/godep)
    * glide [https://github.com/Masterminds/glide](https://github.com/Masterminds/glide)
    * dep [https://github.com/golang/dep](https://github.com/Masterminds/glide)
* glide
    * glide init (初始化，生成glide.yaml配置文件，此文件包含了依赖的包)
    * glide install (安装依赖，会在当前目录下生成vendor目录，同时将第三方依赖下载到vendor目录下)