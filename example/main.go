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

	a := s.OneSelector(`name`)
	b := a.One("first")
	fmt.Println(b)
}
