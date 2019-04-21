# 42 实现micro-kernel framework
## Micro Kernel
### 特点
* 易于扩展
* 错误隔离
* 保证架构一致性
### 要点
* 内核包含公共流程或通用逻辑
* 将可变或可扩展部分规划为扩展点
* 抽象扩展点行为，定义接口
* 利用插件进行扩展
![micro-kernel](micro-kernel.png)
### 示例
![micro-kernel-example](micro-kernel-example.png)