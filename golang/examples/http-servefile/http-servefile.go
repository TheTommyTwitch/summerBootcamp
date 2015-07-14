package main

import (
	"fmt"
	//"io"
	"net/http"
	//"strings"
)


func main() {
	//var dog DogHandler
	//var cat CatHandler

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
	})

	http.HandleFunc("/dogs/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "dog.jpg")
	})

	http.HandleFunc("/cats/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "cat.jpg")
	})

	http.ListenAndServe(":9000", nil)
}
