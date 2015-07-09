package main

import (
  //"strings"
  "fmt"
  "bufio"
  "os"
  "log"
  "io"
)

func Wordcount(scan io.Reader) int {
  word := os.Args[1]
  var count int = 0
  scanner := bufio.NewScanner(scan)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    line := scanner.Text()
    if line == word {
      count++
    }
  }
  return count
}

func main()  {
  srcFile, err := os.Open("book.txt")
	if err != nil {
		log.Fatalln(err)
	}
  defer srcFile.Close()

  fmt.Println(Wordcount(srcFile))
}
