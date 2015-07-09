package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
)

func main() {
	srcFile, err := os.Open("line.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text()
    lineToCap := strings.ToUpper(line[0:1])+line[1:]
		fmt.Println(">>>", lineToCap)
	}
}
