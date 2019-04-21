# 43 内置JSON解析
* 利用反射实现， 通过 FeildTag 来标识对应的json值
```go
    type BasicInfo struct {
    	Name string `json:"name"`
    	Age  int    `josn:"age"`
    }
    type JobInfo struct {
    	Skills []string `json:"skills"`
    }
    type Employee struct {
    	BasicInfo BasicInfo `json:"basic_info"`
    	JobInfo   JobInfo   `josn:"job_info"`
    }
```