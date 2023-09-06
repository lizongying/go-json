# go-re

golang实现的json选择器，使用更简单。

[go-json](https://github.com/lizongying/go-json)

[document](https://pkg.go.dev/github.com/lizongying/go-json)

[english](./README.md)

## 安装

```
go get -u github.com/lizongying/go-json@latest
```

## 用法

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-json/gjson"
)

func main() {
	html := `{"name":{"first":"Janet","last":"Prichard"},"age":47, "a": [2,3]}`
	s, _ := gjson.NewSelectorFromStr(html)

	i := s.One(`age`).Int()
	//47
	fmt.Printf("%+v\n", i)

	u64 := s.One(`age`).Int64()
	//47
	fmt.Printf("%+v\n", u64)

	m := s.Many(`a`)
	n := m[0].Float64()
	//2
	fmt.Println(n)
}

```