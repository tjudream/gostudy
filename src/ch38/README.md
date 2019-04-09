# 38 反射编程
## reflect.TypeOf vs. reflect.ValueOf
* reflect.TypeOf 返回类型(reflect.Type)
* reflect.ValueOf 返回值(reflect.Value)
* 可以从 reflect.Value 获得类型
* 通过 kind 的来判断类型

## 判断类型 - Kind()
```go
    const (
    	Invalid Kind = iota
    	Bool
    	Int
    	Int8
    	Int16
    	Int32
    	Int64
    	Uint
    	Uint8
    	Uint16
    	Uint32
    	Uint64
    ...)
```

## 利用反射编写灵活的代码
* 按名字访问结构的成员
    * reflect.ValueOf(*e).FiledByName("Name")
* 按名字访问结构的方法
    * reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
    
## Struct Tag
```go
    type BasicInfo struct {
    	Name string `json:"name"` //Struct Tag
    	Age  int    `json:"age"`  //Struct Tag
    }
```

## 访问 Struct Tag
```go
    if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
    	t.Error("Failed to get 'Name' field.")
    } else {
    	t.Log("Tag:format", nameField.Tag.Get("format"))
    }
```

* Reflect.Type 和 Reflect.Value 都有FieldByName 方法，注意他们的区别