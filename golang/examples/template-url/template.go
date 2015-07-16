package main

import (
	"log"
	//"os"
	"html/template"
	"net/http"
	"strings"
)

type MyHandler int

type Page struct {
	Title string
	Body  template.HTML
}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error

	url := req.URL.Path

	tpl := template.New("tpl.gohtml")
	tpl = tpl.Funcs(template.FuncMap{
		"uppercase": func(str string) string {
			return strings.ToUpper(str)
		},
	})
	tpl, err = tpl.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(res, Page{
		Title: url,
		Body: `hello world`,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.SetFlags(0)

	var h MyHandler

	http.ListenAndServe(":9000", h)
}
