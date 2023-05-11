package main

import (
	"fmt"
	"practice/Struct/fengzhuang/model"
)

func main() {
	p := model.NewPerson("xining")
	p.SetAge(25)
	fmt.Println(*p)
}