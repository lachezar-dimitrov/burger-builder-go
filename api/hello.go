package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!") // send data to client side
}

func main() {
	http.HandleFunc("/webhook", test)        // set router
	err := http.ListenAndServe(":3030", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
