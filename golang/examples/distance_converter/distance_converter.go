package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
)

const (
  mitoKM = 1.6
  fttoM = 3.28
)

func main() {
  from := os.Args[1]
  to := os.Args[2]

  switch {
  case strings.HasSuffix(from, "mi"):
    switch to {
    case "km":
      miles, _ := strconv.ParseFloat(from[:len(from)-2], 64)
      fmt.Println(miles * mitoKM, "km")
    }
  case strings.HasSuffix(from, "ft"):
    switch to {
    case "m":
      feet, _ := strconv.ParseFloat(from[:len(from)-2], 64)
      fmt.Println(feet * fttoM, "meters")
    }
  case strings.HasSuffix(from, "km"):
    switch to {
    case "mi":
      kilo, _ := strconv.ParseFloat(from[:len(from)-2], 64)
      fmt.Println(kilo / mitoKM, "miles")
    }
  case strings.HasSuffix(from, "m"):
    switch to {
    case "ft":
      meters, _ := strconv.ParseFloat(from[:len(from)-2], 64)
      fmt.Println(meters / mitoKM, "feet")
    }
  }
}
