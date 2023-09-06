# go-json

Json Selector in Golang for Easier Use.

[go-json](https://github.com/lizongying/go-json)

[document](https://pkg.go.dev/github.com/lizongying/go-json)

[中文](./README_CN.md)

## Install

```
go get -u github.com/lizongying/go-json@latest
```

## Usage

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
	//123
	fmt.Printf("%+v\n", i)

	u64 := s.One(`age`).Int64()
	//123
	fmt.Printf("%+v\n", u64)

	m := s.Many(`a`)
	n := m[0].Float64()
	//abc
	fmt.Println(n)
}

```