package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "log"
)


func main()  {
  f, err := os.Open("hello.txt")
  if err != nil {
    log.Fatalln("broke", err.Error())
  }
  defer f.Close()

  bs, err := ioutil.ReadAll(f)
  if err != nil {
    log.Fatalln("broke")
  }

  str := string(bs)
  fmt.Println(str)

  //Writing
  _, err := f.Write(bs)
  if err != nil {
    log.Falalln("broke", err.Error())
  }
}
