## 27 channel的关闭和广播

* channel 的关闭
    * 向关闭的channel 发送数据，会导致 panic
    * v,ok <-ch;ok 为 bool 值，true 表示正常接受，false 表示通道关闭
    * 所有的 channel 接收者都会在 channel 关闭时，立刻从阻塞等待中返回且上述
    ok值为false。这个广播机制常被利用，进行向多个订阅者同时发送信号。
    如：退出信号。