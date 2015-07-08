package main

import "fmt"

func func main() {
  //Array
  var x [4]int
  x[0] = 43
  x[1] = 65
  x[2] = 34
  x[3] = 76

  //Slice
  var y = []float64{1, 2, 3}
  z := y[1:]
  //returns {2, 3}
  fmt.Println(z)

  //append
  xs := []int{1, 2, 3}
  xs = append(xs, 4)

  //copy
  ys := make([]int, 4)
  copy(ys, xs)

  //Maps, similar to js objects
  var t  = make(map[string]int)
  t["key"] = 10
  fmt.Println(x["key"]) //returns 10
  delete(t, "key") //you can delete things
}
