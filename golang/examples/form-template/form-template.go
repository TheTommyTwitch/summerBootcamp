package main

import (
  "net/http"
  "html/template"
  "log"
  "fmt"
)

func main()  {
  log.SetFlags(0)

  http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-type", "text/html")

		firstname := req.FormValue("firstname")
    lastname := req.FormValue("lastname")
    //tpl := template.New("asd")
    tpl, err := template.ParseFiles("tmpl.html")
  	if err != nil {
  		log.Fatalln(err)
  	}
  	err = tpl.Execute(res, map[string]interface{}{
      "firstname": firstname,
      "lastname": lastname,
    })
  	if err != nil {
  		log.Fatalln(err)
  	}
		fmt.Println("value: ", firstname)
    fmt.Println("value: ", lastname)

    tpl = tpl.Funcs(template.FuncMap{
  		"firstname": func() string {
  			return firstname
  		},
      "lastname": func() string {
        return lastname
      },
  	})
  })

  http.ListenAndServe(":9000", nil)
}
