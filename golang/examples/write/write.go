package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "log"
)


func main()  {
  f, err := os.Create("hello.txt")
  if err != nil {
    log.Fatalln("broke", err.Error())
  }
  defer f.Close()

  str := "txt"
  bs := []byt(str)

  _, err := f.Write(bs)
  if err != nil {
    log.Falalln("broke", err.Error())
  }
}
