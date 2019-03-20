## 26 多路选择和超时
* select
    * 与case的顺序无关 
* 多渠道选择
    ```go
    select {
        case ret := <-retCh1:
          t.Logf("result %s", ret) 	 
        case ret := <-retCh2:
          t.Logf("result %s", ret)
        default :
          t.Error("No one returned")
    }
    ```
    
* 超时控制
    ```go
    select {
      case ret := <-retCh:
    	t.Logf("result %s", ret)
      case <-time.After(time.Second * 1):
    	t.Error("time out")
    } 
    ```
