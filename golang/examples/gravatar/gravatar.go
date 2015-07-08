package main

import (
  "fmt"
  "os"
  "io"
  "encoding/hex"
  "strings"
)


func getGravatarHash(email string) string {
  email = strings.TrimSpace(email)
  email = strings.ToLower(email)

  h := md5.New()
  io.WriteString(h, email)
  finalBytes := h.Sum(nil)
  finalString := hex.EncodeToString(finalBytes)
  return finalString
}

func main()  {
  gravatarHash := getGravatarHash("Tasadurian@gmail.com")
  fmt.Fprint(os.Stderr, "Enter your name: ")
  var name string
  fmt.Scanln(&name)
  fmt.Println(`<!DOCTYPE html>
    <html>
      <head>
      <body>
        <p>` + name + `</p>
        <img src="http://www.gravatar.com/avatar/` + gravatarHash + `?d=identicon
      </body>
    </html>`)
}
