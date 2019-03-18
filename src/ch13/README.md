# 18 不一样的接口类型，一样的多态
* Go接口最佳实践：
	1. 倾向于使用小的接口定义，很多接口只包含一个方法
	```
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
		type Write interface {
			Write(p []byte)(n int, err error)
		}
	```
	2. 较大的接口定义，可以由多个小接口定义组合而成
	```
		type ReadWriter interface {
			Reader
			Writer
		}
	```
	3. 只依赖于必要功能的最小接口
	```
		func StoreData(reader Reader) error {
			...
        }
    ```
* 空接口与断言
1. 空接口可以表示任何类型
2. 通过断言来将空接口转换为制定类型
    * v,ok := p.(int) // ok=true 时为转换成功
