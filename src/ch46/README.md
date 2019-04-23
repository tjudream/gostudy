#46 构建Restful服务

## 更好的Router

[https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

```go
    func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
    }

    func main() {
    	router := httprouter.New()
    	router.GET("/", Index)
    	router.GET("/hello/:name", Hello)
    	
    	log.Fatal(http.ListenAndServe(":8080", router))
    }
```

## 面向资源的架构（Resource Oriented Architecture）

In software engineering, a resource-oriented-architecture (ROA) is a style of software architecture and programming 
paradigm for supporitive designing and developing software in the form of Internetworking of resources with "RESTful"
 interfaces.
 
www.demo_portal.com/employee/{name}