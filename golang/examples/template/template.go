package main

import (
	"log"
	//"os"
	"text/template"
	"net/http"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.SetFlags(0)

	var h MyHandler

	http.ListenAndServe(":9000", h)
}
