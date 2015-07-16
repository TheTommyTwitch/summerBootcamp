package main

import (
	"io"
	"net/http"
)
func cCounter(v int) string {
  var output string
  if v > -1 {
    v++
    output = string(v)
  }
  return output
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		// cookie is not set
    cookieCounter := cCounter(1)
		if err != nil {
			//id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-id",
        Value: cookieCounter,
			}
		}

		http.SetCookie(res, cookie)

		io.WriteString(res, `<!DOCTYPE html>
<html>
  <body>
    <h1>`+cookie.Value+`</h1>
  </body>
</html>`)

	})
	http.ListenAndServe(":9000", nil)
}
