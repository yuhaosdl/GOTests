package main

import (
	"fmt"
)

type Foo struct {
	Name string
	Age  int
}

func (self *Foo) Get() {
	self.Name = "测试"
	return
}

func (self Foo) GetConst() {
	self.Name = "测试"
	return
}
func main() {
	fmt.Println(2123123)
	people := People{Name: "yuhao", Age: 24}
	foo := Foo{Name: "校长", Age: 12}
	fmt.Println(people)
	alter(people)
	fmt.Println(people)
	var a Fee
	a = &people
	fmt.Printf("%p", &people)
	fmt.Println()
	people.Feed()
	fmt.Println()
	a.Feed()
	tttt(a)
	fmt.Println(foo)
	foo.Get()
	fmt.Println(foo)
}
func alter(people People) {
	people.Name = "Lili"
}

type People struct {
	Name string
	Age  int
}

func (people *People) Feed() {
	fmt.Printf("接口测试啊啊啊啊 %p", &people)
	fmt.Println()
}

type Fee interface {
	Feed()
}

func tttt(fee Fee) {
	fee.Feed()
}
