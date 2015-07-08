package main

import "fmt"

func main()  {
  numbers := increasingNumberGenerator()
  fmt.Println(increment())
  fmt.Println(increment())

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}

func increasingNumberGenerator() func() int {
  x := 0
    return func() int {
      x++
      return x
    }
}

func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
