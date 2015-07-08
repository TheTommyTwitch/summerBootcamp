package main

import "fmt"
import "strconv"
import "os"

const mitoKM = 1.6

func main() {
  number, err := strconv.ParseFloat(os.Args[1], 64)
  if err != nil {
    panic(err)
  }
  fmt.Println("<!DOCTYPE html>")
  fmt.Println("<html>")
  fmt.Println(" <head></head>")
  fmt.Println(" <body>")

  fmt.Println("Miles: ", number)
  fmt.Println("KM: ", (number * mitoKM))

  fmt.Println(" </body>")
  fmt.Println("</html")

}
