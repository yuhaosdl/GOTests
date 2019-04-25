package main

import (
	"errors"
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

var arr [36]string = [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var temp = make(map[string]int)

func main() {
	result := Calculate("1z", "2", "4")
	fmt.Println(result)
}

// 计算
func Calculate(args ...string) (result string) {
	t := 0
	for _, v := range args {
		//转成10进制
		if v2, err := convertToTen(v); err != nil {
			panic(err.Error())
		} else {
			t = t + v2
		}
		//加和

	}
	//转回36进制
	result = convertToThirtySix(t)
	return
}

//转换成36进制
func convertToThirtySix(in int) (out string) {
	a := in / (len(arr))
	out = arr[in%(len(arr))]
	if a != 0 {
		out = convertToThirtySix(a) + out
	}
	return
}

//转换成10进制
func convertToTen(str string) (result int, err error) {
	//先查缓存
	if v, ok := temp[str]; ok {
		result = v
	} else {
		for _, c := range str {
			if v, ok := temp[string(c)]; ok {
				result = 10*result + v
			} else {
				var flag bool
				for i, v2 := range arr {
					if v2 == string(c) {
						temp[v2] = i
						result = len(arr)*result + i
						flag = true
						break
					}
				}
				if !flag {
					err = errors.New("not found!")
					break
				}
			}
		}
	}
	return
}

func ChangePeople(p *People) {
	fmt.Printf("ChangePeople中p的地址是%p\n", p)
	fmt.Println(&p)
	fmt.Printf("ChangePeople中p栈上的地址是%p\n", &p)
	p = &People{
		Age:  12,
		Name: "sw",
	}
	//fmt.Printf("ChangePeople中p的地址是%p\n", p)
}
func AlterName(p *People) {
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
