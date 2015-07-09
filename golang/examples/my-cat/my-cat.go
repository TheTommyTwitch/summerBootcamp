package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "log"
)


func main()  {
  file := os.Args[1]
  f, err := os.Open(file)
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
}
