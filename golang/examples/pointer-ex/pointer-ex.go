//Write a program that can swap two integers

package main

import "fmt"

func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func rotate(args ...*int)  {
  if len(args) == 0 {
    return
  }
  firstValue := *args[0]
  for i := 0; i <= len(args); i++ {
    *ptr = range *args[i + 1]
  }
  *args[len(args)-1] = firstValue
}

func main()  {
  x := 2
  y := 5
  swap(&x, &y)
  fmt.Println(x, y)
}
