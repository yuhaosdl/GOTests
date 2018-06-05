package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func (people *People) Feed() {
	fmt.Println("接口测试啊啊啊啊")
}

type Fee interface {
	Feed()
}

func tttt(fee Fee) {
	fee.Feed()
}

func (people *People) test() {

}
