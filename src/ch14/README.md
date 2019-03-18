# 19 编写好的错误处理
* Go的错误机制
    * 与其他主要编程语言的差异
        1. 没有异常机制
        2. error类型实现了error接口
        3. 可以通过 errors.New 来快速创建错误实例
        ```go
              type error interface {
                  Error() string
              }
              errors.New("n must be in the range [0,10]")
        ```
* 定义自己的错误类型:         
var LessThanTwoError = errors.New("n should be not less than 2")
* 最佳实践：及早失败，避免嵌套

# <h2 id="1"> 20 panic和recover </h2>
* panic 
    * panic 用于不可以恢复的错误
    * panic 退出前会执行 defer 指定的内容
* panic vs os.Exit
    * os.Exit 退出时不会调用 defer 指定的函数
    * os.Exit 退出时不输出当前调用栈信息
* recover
    ```go
      defer func() {
    	if err := recover(); err != nil {
  		    //恢复错误
  	    }
      }()
    ```
* 当心 recover 成为恶魔
    * 形成僵尸服务进程， 导致 health check 失效
    * "Let it Crash!" 往往是我们恢复不确定性错误的最好方法
    
