package main

import (
  //"fmt"
  "io/ioutil"
  "os"
  "log"
)
func mutireader() string {
  var str string
  for i, v := range os.Args[1:] {
    f, err := os.Open(v)
    if err != nil {
      log.Fatalln("broke", err.Error())
    }
    defer v.Close()
    

    io.Copy(os.Stdout, f)
}

func main()  {
  mutireader()
}
