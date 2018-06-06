package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseHandler, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandlerFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
