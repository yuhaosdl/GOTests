package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ceshi
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "<h1>Hello World!</h1>")
	})
	http.ListenAndServe(":8888", nil)
}
