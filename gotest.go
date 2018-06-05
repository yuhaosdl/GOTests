package main

import (
	"fmt"
)

func main() {
	fmt.Println(2123123)
	people := People{Name: "yuhao", Age: 24}
	fmt.Println(people)
	alter(people)
	fmt.Println(people)
	var a Fee
	a = &people
	a.Feed()
	tttt(a)
}
func alter(people People) {
	people.Name = "Lili"
}
