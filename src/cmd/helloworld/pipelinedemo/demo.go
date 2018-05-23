package main

import (
	"fmt"
	"pipeline"
)

func main() {
	a := pipeline.DataSort(pipeline.DataSource(9, 3, 4, 5, 2, 9, 1, 8, 3, 0, 4, 8, 3, 6, 2, 5, 1, 0, 4, 8, 3, 1))
	for v := range a {
		fmt.Println(v)
	}
}
