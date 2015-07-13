package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
//red
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	r, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(r)

//blue
  type Blue struct {
    Colors []string
  }
  blueGroup := Blue{
    Colors: []string{"Blue", "Sky", "Baby"},
  }
  b, err := json.Marshal(blueGroup)
  if err != nil {
    fmt.Println("error", err)
  }
  os.Stdout.Write(b)
}
