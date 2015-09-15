package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(rep http.ResponseWriter, req *http.Request) {
		rep.Write([]byte("hello"))
	})
	fmt.Println("listening  5642")
	http.ListenAndServe(":5642", nil)	
}
