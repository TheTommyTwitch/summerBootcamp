package main

import (
  "strings"
  "fmt"
)

func Wordcount(str string) map[string]int {
  words := strings.Fields(str)
  var thisMap = make(map[string]int)
  var count int = 0
  for i := 0; i < len(words); i++ {
    count++
    thisMap[words[i]] = count
  }
  return thisMap
}

func main()  {
  str := "Whatever I want to hear in my head"
  result := Wordcount(str)
  fmt.Println(result)

}
