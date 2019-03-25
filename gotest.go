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

	p := People{
		Age:  12,
		Name: "yuhao",
	}
	AlterName(p)
	fmt.Println(p.Name)

}
func AlterName(p People) {
	p.Name = "yyyy"
}

type TestChanel struct {
	inChanel chan string
	close    chan string
}

func alter(people People) {
	people.Name = "Lili"
}

type People struct {
	Name string
	Age  int
}

func (people *People) Feed() {
	fmt.Printf("接口测试改动 %p", &people)
	fmt.Println()
}

type Fee interface {
	Feed()
}

func tttt(fee Fee) {
	fee.Feed()
}
