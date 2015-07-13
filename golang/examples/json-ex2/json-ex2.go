package main

import (
  "fmt"
  "encoding/json"
  "os"
)

type myData struct {
  ID int
  Name string
  Price float64
  Tags []string
}

var x myData

func main()  {
  data, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer data.Close()

    enc := json.NewDecoder(data)
    enc.Decode(&x)
    fmt.Println(x)
}
