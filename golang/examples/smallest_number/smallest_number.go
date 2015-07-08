package main

import "fmt"

func main() {
even, h := halves(5)
fmt.Println(even, h)

gs := []int{
   48,96,86,68,
   57,82,63,70,
   37,34,83,27,
   19,97, 9,17,
}
g := greatestNumber(gs)
fmt.Println(g)

}

//finds the smallest number in an array
func smallest() {
  xs := []int{
     48,96,86,68,
     57,82,63,70,
     37,34,83,27,
     19,97, 9,17,
  }
  min := xs[0]
  for _, value := range xs {
    if value < min {
      min = value
    }
  }
  fmt.Println(min)
}

//returns the sum
func sum(x []float64) float64 {
  total := 0.0
  for _, v := range x {
    total += v
  }
  return total
}

//returns half the inputed number and a boolean for odd or even
func halves(y int) (bool, int) {
  var oddOrEven bool = false
  var dividedNumber int = y / 2
  if y % 2 == 0 {
    oddOrEven = true
  } else {
    oddOrEven = false
  }
  return oddOrEven, dividedNumber
}

//Write a function with one variadic parameter that
//finds the greatest number in a list of numbers.

func greatestNumber(x[]int) int {
  great := x[0]
  for _, value := range x {
    if value > great {
      great = value
    }
  }
  return great
}

//Using makeEvenGenerator as an example,
//write a makeOddGenerator function that generates odd numbers.

func makeOddGenerator() {
  i := int(1)
  func() {

  }
}

//The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1,
//fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).

func fib(n int) int {
  if n == 0 {
    return 0;
  } else if n == 1 {
    return 1;
  } else {
    return n = (n-1) + (n-2)
  }
}

//Create a function which reverses a list of integers:
//reverse([]int{1,2,3}) → []int{3,2,1}

func reverse(x[]int) []int {
  
}
