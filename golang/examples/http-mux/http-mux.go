package main

import (
	"io"
	"net/http"
)

type DogHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-type", "text/html")
	io.WriteString(res, `<img src="http://www.globalpetfoodskw.ca/images/dog_img.png">`)
}

type CatHandler int

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-type", "text/html")
	io.WriteString(res, `<img src="http://www.cats.org.uk/uploads/images/pages/photo_latest14.jpg">`)
}

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":9000", mux)
}
