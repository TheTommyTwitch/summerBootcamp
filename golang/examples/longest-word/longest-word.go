package main

import (
  //"strings"
  "fmt"
  "bufio"
  "os"
  "log"
  "io"
)

func Wordcount(scan io.Reader) string {
  scanner := bufio.NewScanner(scan)
  scanner.Split(bufio.ScanWords)
  var longestWord string
  for scanner.Scan() {
    line := scanner.Text()
    wordLength := len(line)
    if wordLength > len(longestWord) {
      longestWord = line
    }
  }
  return longestWord
}

func main()  {
  srcFile, err := os.Open("book.txt")
	if err != nil {
		log.Fatalln(err)
	}
  defer srcFile.Close()

  fmt.Println(Wordcount(srcFile))
}
