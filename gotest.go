package main

import (
	"fmt"
	"time"
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
	ch := make(chan string)
	go func() {
		for {
			ch <- "11"
		}

	}()
	time.Sleep(1 * time.Second)
	fmt.Println(<-ch)
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
	fmt.Printf("接口测试啊啊啊啊 %p", &people)
	fmt.Println()
}

type Fee interface {
	Feed()
}

func tttt(fee Fee) {
	fee.Feed()
}
