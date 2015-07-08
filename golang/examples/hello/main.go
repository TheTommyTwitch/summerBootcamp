package main

import "fmt"

const mtoKM = 1.6
const ftoC = -17.2222
const ptoKG = 0.45

func main() {
  // var name string
  // fmt Print("Enter your name")
  // fmt.ScanF("%s" + &name)
  // fmt.Println()
  var miles float64
  fmt.Print("Enter miles: ")
  fmt.Scanf("%v", &miles)
  kilometers := (miles * mtoKM)
  fmt.Println("---------------")
  fmt.Printf("| Miles: %10.2f |\n", miles)
  fmt.Println("---------------")
  fmt.Printf("| Kilometers: %10.2f |\n", kilometers)
  fmt.Println("---------------")

  // var degrees float64
  // fmt.Print("Enter temerature in degrees: ")
  // fmt.Scanf("%v", &degrees)
  // cel := (degrees - 32 * (5.0 / 9))
  // fmt.Println("Temperature in Celcius: ", cel)
  //
  // var pounds float64
  // fmt.Print("Enter pounds: ")
  // fmt.Scanf("%v", &pounds)
  // kg := (pounds * ptoKG)
  // fmt.Println("This weight in kilograms is: ", kg)
}
