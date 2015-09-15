package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(rep http.ResponseWriter, req *http.Request) {
		rep.Write([]byte("hello"))
	})
	fmt.Println("listening  5643")
	http.ListenAndServe(":5643", nil)	
}
