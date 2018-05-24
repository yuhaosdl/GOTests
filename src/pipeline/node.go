package pipeline

import (
	"sort"
)

//DataSource input into  channel
func DataSource(in ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range in {
			out <- v
		}
		close(out)
	}()
	return out
}

//DataSort channel Sort
func DataSort(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		arr := []int{}
		for v := range in {
			arr = append(arr, v)
		}
		sort.Ints(arr)
		for _, v := range arr {
			out <- v
		}
		close(out)
	}()
	return out
}
