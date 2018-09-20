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
	testChanel := &TestChanel{
		inChanel: make(chan string, 10),
		close:    make(chan string, 1),
	}
	go func() {
		time.Sleep(5 * time.Second)
		close(testChanel.close)
	}()
	for {
		select {
		case testChanel.inChanel <- "1111":
			fmt.Print("没关闭")
		case a := <-testChanel.close:
			fmt.Println(a)
		}
	}

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
