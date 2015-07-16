package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type index struct {
	Photos []string
	Body   string
}

func myFiles() []string {
	files := make([]string, 0)
	filepath.Walk("img", func(p string, info os.FileInfo, _ error) error {
		if strings.HasSuffix(p, ".jpeg") || strings.HasSuffix(p, ".jpg") {
			files = append(files, p)
		}
		return nil
	})
	return files
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		tpl := template.New("index.tpl")
		tpl = tpl.Funcs(template.FuncMap{
		//function go here {},
		})
		tpl, err := tpl.ParseFiles("index.tpl")
		if err != nil {
			log.Fatalln(err)
		}
		err = tpl.Execute(res, index{
			Photos: myFiles(),
		})
	})

	// http.HandleDunc("/admin", func(res http.ResponeWriter, req *http.Request) {
	//
	// })

	http.ListenAndServe(":8080", nil)
}
