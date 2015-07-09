package main

import (
  //"fmt"
  "io/ioutil"
  "os"
  "log"
)

//Read File
func main()  {
  ReadFile := os.Args[1]
  WriteFile := os.Args[2]

  f, err := os.Open(ReadFile)
  if err != nil {
    log.Fatalln("broke", err.Error())
  }
  defer f.Close()

  //Write File
  w, err := os.Create(WriteFile)
  if err != nil {
    log.Fatalln("broke", err.Error())
  }
  defer w.Close()

  _, err = io.Copy(w, f)
  if err != nil {
    return fmt.Errorf("error copying file: %v", err)
  }
  return nil
}
