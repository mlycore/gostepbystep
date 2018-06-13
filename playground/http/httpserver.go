package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!")
}

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenandServer error:", err.Error())
	}
}

